package config

//配置实体
type Config struct {
	LogLevel   string //日志级别: debug, info, warning, error, fatal, panic
	MongoDBHosts []string //mongo数据库ip+port
}

//type Redis struct {
//	NetType  string `toml:"netType"`
//	Addr     string `toml:"addr"`
//	Password string `toml:"password"`
//}
//
//type DB struct {
//	Host     string       `json:"host"`
//	Port     string        `json:"port"`
//	User     string       `json:"user"`
//	Password string   `json:"password"`
//	DbName   string     `json:"dbName"`
//}
//
//type Mongodb struct {
//	Dsn      string `json:"dsn"`
//	Password string   `json:"password"`
//	DbName   string     `json:"dbName"`
//}
