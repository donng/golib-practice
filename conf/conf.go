package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Database struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string `mapstructure:"database"` // 结构体命名与配置文件不一致，需要声明映射关系
}

type Redis struct {
	Host     string
	Port     int
	UserName string
	Password string
}

type Conf struct {
	Database Database `mapstructure:"mysql"`
	Redis    Redis
}

var Setting = &Conf{}

func init() {
	viper.SetConfigName("conf") // 声明文件名
	viper.SetConfigType("toml") // 声明文件类型
	viper.AddConfigPath(".")    // 声明文件路径

	// 读取文件配置
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	// 解析结构体
	if err := viper.Unmarshal(Setting); err != nil {
		log.Fatal(err)
	}
}

func ShowViperUsage() {

}
