package conf

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
)

type server struct {
	HttpPort  string `mapstructure:"httpPort"`
	JWTSecret string `mapstructure:"jwtSecret"`
	JWTExpire int    `mapstructure:"jwtExpire"`
	DomainName string `mapstructure:"domainName"`
}

var ServerConf = &server{}

type database struct {
	DBType   string `mapstructure:"dbType"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbName"`
	Debug    bool   `mapstructure:"debug"`
}

var DBConf = &database{}

// 初始化配置文件
func Setup() {
	viper.SetConfigType("yaml")
	confFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("读取配置文件失败")
	}
	err = viper.ReadConfig(bytes.NewBuffer(confFile))
	if err != nil {
		fmt.Println("Viper库读取配置文件失败")
	}
	viper.UnmarshalKey("server", ServerConf)
	viper.UnmarshalKey("database", DBConf)
}
