package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config contains server configurations.
type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfig
}

// ServerConfig holds configuration specific to server instantiation.
type ServerConfig struct {
	Port         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
}

// DatabaseConfig for database configuration.
type DatabaseConfig struct {
	Port     int
	Client   string
	Host     string
	User     string
	Password string
	Db       string
}

// Initialize env variables from env file.
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

// New configurations for server.
func New() (*Config, error) {

	var c Config

	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	server := &ServerConfig{
		Port:         c.Server.Port,
		ReadTimeOut:  c.Server.ReadTimeOut * time.Second,
		WriteTimeOut: c.Server.WriteTimeOut * time.Second,
	}

	db := &DatabaseConfig{
		Port:     c.Database.Port,
		User:     c.Database.User,
		Password: c.Database.Password,
		Client:   c.Database.Client,
		Host:     c.Database.Host,
		Db:       c.Database.Db,
	}

	return &Config{
		Server:   server,
		Database: db,
	}, nil
}
