package config


//配置实体
type Config struct {
	//日志级别: debug, info, warning, error, fatal, panic
	LogLevel string
	//mongo数据库配置
	MongoDBHosts []string
}

