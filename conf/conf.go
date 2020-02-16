package conf

import (
	"log"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Port     int
	Host     string
	Username string
	Password string
	Database string
}

var Mysql MysqlConfig

func init() {
	viper.SetConfigName("conf")                     // 读取yaml配置文件
	viper.AddConfigPath("$GOPATH/src/qasite/conf/") //设置配置文件的搜索目录
	err := viper.ReadInConfig()                     // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	Mysql.Port = viper.GetInt("mysql.port")
	Mysql.Host = viper.GetString("mysql.host")
	Mysql.Username = viper.GetString("mysql.username")
	Mysql.Password = viper.GetString("mysql.password")
	Mysql.Database = viper.GetString("mysql.database")
}
