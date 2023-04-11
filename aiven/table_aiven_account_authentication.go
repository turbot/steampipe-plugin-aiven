package aiven

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableAivenAccountAuthentication(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aiven_account_authentication",
		Description: "Retrieve information about your account authentications.",
		List: &plugin.ListConfig{
			ParentHydrate: listAccounts,
			Hydrate:       listAccountAuthentications,
			KeyColumns: []*plugin.KeyColumn{
				{
					Name:    "account_id",
					Require: plugin.Optional,
				},
			},
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"account_id", "id"}),
			Hydrate:    getAccountAuthentication,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "Authentication method ID.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Authentication method name.",
			},
			{
				Name:        "account_id",
				Type:        proto.ColumnType_STRING,
				Description: "The account ID.",
			},
			{
				Name:        "enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "If true, authentication method can be used to access account/projects in account. If false, authentication method can still be used to sign in.",
			},
			{
				Name:        "state",
				Type:        proto.ColumnType_STRING,
				Description: "The state of the authentication method.",
			},
			{
				Name:        "auto_join_team_id",
				Type:        proto.ColumnType_STRING,
				Description: "Automatically add users to a team, when user signs up using this authentication method.",
			},
			{
				Name:        "create_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The create time of the authentication method.",
			},
			{
				Name:        "delete_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The delete time of the authentication method.",
			},
			{
				Name:        "saml_acs_url",
				Type:        proto.ColumnType_STRING,
				Description: "Saml acs url.",
				Transform:   transform.FromField("SAMLAcsUrl"),
			},
			{
				Name:        "saml_certificate",
				Type:        proto.ColumnType_STRING,
				Description: "Identity provider's certificate.",
				Transform:   transform.FromField("SAMLCertificate"),
			},
			{
				Name:        "saml_certificate_issuer",
				Type:        proto.ColumnType_STRING,
				Description: "Saml certificate issuer.",
				Transform:   transform.FromField("SAMLCertificateIssuer"),
			},
			{
				Name:        "saml_certificate_not_valid_after",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Saml certificate not valid after.",
				Transform:   transform.FromField("SAMLCertificateNotValidAfter"),
			},
			{
				Name:        "saml_certificate_not_valid_before",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Saml certificate not valid before.",
				Transform:   transform.FromField("SAMLCertificateNotValidBefore"),
			},
			{
				Name:        "saml_certificate_subject",
				Type:        proto.ColumnType_STRING,
				Description: "Saml certificate subject.",
				Transform:   transform.FromField("SAMLCertificateSubject"),
			},
			{
				Name:        "saml_entity",
				Type:        proto.ColumnType_STRING,
				Description: "Saml entity.",
				Transform:   transform.FromField("SAMLEntity"),
			},
			{
				Name:        "saml_idp_url",
				Type:        proto.ColumnType_STRING,
				Description: "Saml idp url.",
				Transform:   transform.FromField("SAMLIdpUrl"),
			},
			{
				Name:        "saml_metadata_url",
				Type:        proto.ColumnType_STRING,
				Description: "Saml metadata url.",
				Transform:   transform.FromField("SAMLMetadataUrl"),
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Authentication method type.",
			},
			{
				Name:        "update_time",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The update time of the authentication method.",
			},
		},
	}
}

func listAccountAuthentications(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	account := h.Item.(aiven.Account)
	account_id := d.EqualsQuals["account_id"].GetStringValue()

	// check if the provided account_id is not matching with the parentHydrate
	if account_id != "" && account_id != account.Id {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_account_authentication.listAccountAuthentications", "connection_error", err)
		return nil, err
	}

	authenticationList, err := conn.AccountAuthentications.List(account.Id)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_account_authentication.listAccountAuthentications", "api_error", err)
		return nil, err
	}

	for _, authenticationMethod := range authenticationList.AuthenticationMethods {
		d.StreamListItem(ctx, authenticationMethod)

		// Context can be cancelled due to manual cancellation or the limit has been hit
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func getAccountAuthentication(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	account_id := d.EqualsQuals["account_id"].GetStringValue()
	auth_id := d.EqualsQuals["id"].GetStringValue()

	// Check if account_id or auth_id is empty
	if account_id == "" || auth_id == "" {
		return nil, nil
	}

	conn, err := getClient(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_account_authentication.getAccountAuthentication", "connection_error", err)
		return nil, err
	}

	accountAuthentication, err := conn.AccountAuthentications.Get(account_id, auth_id)
	if err != nil {
		plugin.Logger(ctx).Error("aiven_account_authentication.getAccountAuthentication", "api_error", err)
		return nil, err
	}

	return accountAuthentication.AuthenticationMethod, nil
}
