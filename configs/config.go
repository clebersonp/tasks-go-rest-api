// Manage all settings for Rest API and database configuration like database port, API port and so on.
// Everything will be loaded from file using viper third-party package. https://pkg.go.dev/github.com/spf13/viper
package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API apiConfig
	DB  dbConfig
}

type apiConfig struct {
	Port string // port of server rest api
}

type dbConfig struct {
	Host     string // db host name
	Port     string // db port number
	User     string // db user name
	Password string // db password
	Database string // db name
}

// It is always called at application start automatically
// It uses the third party viper package to manager the environment variables of files
func init() {
	// If any flags are not provided from env files and instead use default values
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// Defines the viper configuration and loads all values ​​from the provided configuration file
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// if you didn't find the file, that's ok, as we had already defined some default values
			return err
		}
	}
	// same as &config{}
	cfg = new(config)

	// Load api-config values
	cfg.API = apiConfig{
		Port: viper.GetString("api.port"),
	}

	// Load db-config values
	cfg.DB = dbConfig{
		Host: viper.GetString("database.host"),
		Port: viper.GetString("database.port"),
		User: viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),	
	}

	return nil
}

// DB returns the DB configurations
func DB() dbConfig {
	return cfg.DB
}

// ServerPort returns the API port configuration
func ServerPort() string {
	return cfg.API.Port
}