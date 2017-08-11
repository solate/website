package logs

import (
	"github.com/solate/website/sys/config"
	"github.com/Sirupsen/logrus"
	"os"

)

func Init() (error) {
	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.TextFormatter{})

	//设置最低loglevel
	logrus.SetLevel(getLevel(config.Conf.LogLevel))

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
		return err
	}
	return nil
}

func getLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warning":
		return logrus.WarnLevel
	case "error":
		return  logrus.ErrorLevel
	case "fatal":
		return  logrus.FatalLevel
	case "panic" :
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel //默认是info

	}
}
