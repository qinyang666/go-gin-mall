package config

type DBConf struct {
	Host     string
	Port     int32
	Username string
	Password string
	DBName   string
	Debug    bool
}
