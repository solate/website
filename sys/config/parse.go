// 该文件主要是为了封装config
package config

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
	"github.com/solate/website/libs/file"
)

const (
	CONFIGPATH = "conf/config.toml"   //配置文件默认位置
)

var (
	Conf      *Config  // 配置
	configPath string   // 配置文件路径
)

//加载配置
func LoadConfig() (err error) {
	flag.StringVar(&configPath, "c", file.GetCurrentDirectory()+"/"+ CONFIGPATH, "set config file path")
	flag.Parse()
	if _,err = toml.DecodeFile(configPath, Conf); err != nil {
		logrus.Error("failed to load conf/config.toml" + err.Error())
		return
	}
	logrus.Debug("finish Load Config")
	return
}
