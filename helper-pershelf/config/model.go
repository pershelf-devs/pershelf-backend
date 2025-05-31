package config

import (
	config "github.com/pershelf/pershelf/config/database"
	config2 "github.com/pershelf/pershelf/config/server"
)

// DBConfig a struct to use in decoding the database configurations
type Config struct {
	Conn config.DBConnectionConfig `json:"connection"`
	Srv  config2.ServerConfig    `json:"server"`
}
