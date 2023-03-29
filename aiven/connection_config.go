package aiven

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	aivenClient "github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type aivenConfig struct {
	APIKey    *string `cty:"api_key"`
	UserAgent *string `cty:"user_agent"`
	Email     *string `cty:"email"`
	Password  *string `cty:"password"`
	OTP       *string `cty:"otp"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key": {
		Type: schema.TypeString,
	},
	"user_agent": {
		Type: schema.TypeString,
	},
	"email": {
		Type: schema.TypeString,
	},
	"password": {
		Type: schema.TypeString,
	},
	"otp": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &aivenConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) aivenConfig {
	if connection == nil || connection.Config == nil {
		return aivenConfig{}
	}
	config, _ := connection.Config.(aivenConfig)
	return config
}

func getClient(ctx context.Context, d *plugin.QueryData) (*aivenClient.Client, error) {
	aivenConfig := GetConfig(d.Connection)

	apiKey := os.Getenv("AIVEN_TOKEN")
	userAgent := ""
	email := ""
	password := ""
	otp := ""

	if aivenConfig.APIKey != nil {
		apiKey = *aivenConfig.APIKey
	}
	if aivenConfig.Email != nil {
		email = *aivenConfig.Email
	}
	if aivenConfig.Password != nil {
		password = *aivenConfig.Password
	}
	if aivenConfig.OTP != nil {
		otp = *aivenConfig.OTP
	}
	if aivenConfig.UserAgent != nil {
		userAgent = *aivenConfig.UserAgent
	}

	// Authenticate with MFAUser
	if email != "" && password != "" && otp != "" {
		client, err := aivenClient.NewMFAUserClient(email, otp, password, userAgent)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	// Authenticate with User
	if email != "" && password != "" {
		client, err := aivenClient.NewUserClient(email, password, userAgent)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	// Authenticate with API Key
	if apiKey != "" {
		client, err := aivenClient.NewTokenClient(apiKey, userAgent)
		if err != nil {
			return nil, err
		}
		return client, nil
	} else {
		// Authenticate with CLI
		home, _ := os.UserHomeDir()
		file, _ := ioutil.ReadFile(home + "/.config/aiven/aiven-credentials.json")

		cliCreds := make(map[string]string)
		_ = json.Unmarshal([]byte(file), &cliCreds)

		for k, v := range cliCreds {
			if k == "auth_token" {
				apiKey = v
			}
		}

		if apiKey != "" {
			client, err := aivenClient.NewTokenClient(apiKey, userAgent)
			if err != nil {
				return nil, err
			}
			return client, nil
		}
	}

	return nil, errors.New("'api_key' or ('email' and 'password') or ('email', 'password' and 'otp') must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
}
