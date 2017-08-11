package config

import "time"

////配置实体
//type Config struct {
//	LogLevel   string //日志级别: debug, info, warning, error, fatal, panic
//	MongoDBHosts []string //mongo数据库ip+port
//}

//配置实体
//
type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time // requires `import time`

	DB      DB
	Redis   Redis
	Mongodb Mongodb

	//日志级别: debug, info, warning, error, fatal, panic
	LogLevel string
	//mongo数据库配置
	MongoDBHosts []string
}

type Redis struct {
	NetType  string `toml:"netType"`
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
}

type DB struct {
	Host     string       `json:"host"`
	Port     string        `json:"port"`
	User     string       `json:"user"`
	Password string   `json:"password"`
	DbName   string     `json:"dbName"`
}

type Mongodb struct {
	Dsn      string `json:"dsn"`
	Password string   `json:"password"`
	DbName   string     `json:"dbName"`
}