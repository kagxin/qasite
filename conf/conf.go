package conf

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Config struct {
	Mysql *MysqlConfig
}

type MysqlConfig struct {
	Port     int
	Host     string
	Username string
	Password string
	Database string
}

func Init() *Config {
	viper.SetConfigName("conf")                     // 读取yaml配置文件
	viper.AddConfigPath("$GOPATH/src/qasite/conf/") //设置配置文件的搜索目录
	err := viper.ReadInConfig()                     // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	mysql := &MysqlConfig{}
	mysql.Port = viper.GetInt("mysql.port")
	mysql.Host = viper.GetString("mysql.host")
	mysql.Username = viper.GetString("mysql.username")
	mysql.Password = viper.GetString("mysql.password")
	mysql.Database = viper.GetString("mysql.database")
	return &Config{
		Mysql: mysql,
	}
}
