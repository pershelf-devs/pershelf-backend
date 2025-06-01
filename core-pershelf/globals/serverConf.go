package globals

type ServerConfig struct {
	ServerIP   string `json:"serverIp"`
	ServerPort string `json:"serverPort"`
	HelperPort string `json:"helperPort"`
}

type Server struct {
	Server ServerConfig
}

var (
	ServerConf Server
)
