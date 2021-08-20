package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 使用结构体变量报错配置信息
type Config struct {
	*AppConfig   `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Version     string `mapstructure:"version"`
	Port        int    `mapstructure:"port"`
	MachineRoom string `mapstructure:"machine_room"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poll_size"`
}

// 全局变量,用来保存程序的所有配置信息
var Conf = new(Config)

func Init(filename string) (err error) {
	viper.SetConfigFile(filename) //指定配置文件路径

	//viper.SetConfigName("config") //配置文件名称
	//viper.AddConfigPath("./")     //配置文件路径,可以添加多个

	//viper.SetConfigType("yaml")   //远程加载配置时,指定配置的类型

	err = viper.ReadInConfig() //查找并加载配置文件
	if err != nil {            //处理加载配置文件的错误
		fmt.Printf("viper read config failed, err:%v\n", err)
		return
	}
	//把读取到的配置信息反序列化到结构体变量 Conf 中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("vipre unmarshal failed, err:%v\n", err)
		return
	}

	viper.WatchConfig() //监听配置变化
	//配置发生变化时的处理函数(把新配置重新加载到 Conf 变量)
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("the config has changed!")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("vipre unmarshal failed, err:%v\n", err)
			return
		}
	})
	return
}
