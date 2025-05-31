package config

// DBConnectionConfig a struct to use in decoding the database-connection configurations
type DBConnectionConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
	Network  string `json:"network"`
}
