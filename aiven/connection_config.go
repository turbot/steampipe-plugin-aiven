package aiven

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	aivenClient "github.com/aiven/aiven-go-client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type aivenConfig struct {
	APIKey   *string `hcl:"api_key"`
	Email    *string `hcl:"email"`
	Password *string `hcl:"password"`
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
	email := ""
	password := ""

	if aivenConfig.APIKey != nil {
		apiKey = *aivenConfig.APIKey
	}
	if aivenConfig.Email != nil {
		email = *aivenConfig.Email
	}
	if aivenConfig.Password != nil {
		password = *aivenConfig.Password
	}

	if apiKey != "" { // Authenticate with API Key
		client, err := aivenClient.NewTokenClient(apiKey, "")
		if err != nil {
			return nil, err
		}
		return client, nil
	} else if email != "" && password != "" { // Authenticate with User
		client, err := aivenClient.NewUserClient(email, password, "")
		if err != nil {
			return nil, err
		}
		return client, nil
	} else { // Authenticate with CLI
		home, _ := os.UserHomeDir()
		file, _ := os.ReadFile(home + "/.config/aiven/aiven-credentials.json")

		cliCreds := make(map[string]string)
		_ = json.Unmarshal([]byte(file), &cliCreds)

		for k, v := range cliCreds {
			if k == "auth_token" {
				apiKey = v
			}
		}

		if apiKey != "" {
			client, err := aivenClient.NewTokenClient(apiKey, "")
			if err != nil {
				return nil, err
			}
			return client, nil
		}
	}

	return nil, errors.New("'api_key' or ('email' and 'password') must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe.")
}
