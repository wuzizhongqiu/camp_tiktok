package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var globalConfig = new(GlobalConfig)

type GlobalConfig struct {
	*SvrConfig    `mapstructure:"svr_config"`
	*RedisConfig  `mapstructure:"redis"`
	RedSyncConfig []RedSyncConfig `mapstructure:"redsync"`
	*DBConfig     `mapstructure:"mysql"`
	*LogConfig    `mapstructure:"log"`
	*ConsulConfig `mapstructure:"consul"`
}

type SvrConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructuer:"port"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	PassWord string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	Port     int    `mapstructure:"port"`
	PoolSize int    `mapstructure:"PoolSize"`
}

type RedSyncConfig struct {
	Host       string `mapstructure:"host"`
	PassWord   string `mapstructure:"password"`
	Port       int    `mapstructure:"port"`
	LockExpire int    `mapstructure:"LockExpire"`
	PoolSize   int    `mapstructure:"PoolSize"`
}

type DBConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	DataBase    string `mapstructure:"database"`
	UserName    string `mapstructure:"username"`
	PassWord    string `mapstructure:"password"`
	MaxIdleConn int    `mapstructure:"max_id_cnn"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleTime int    `mapstructure:"max_idle_time"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"file_name"`
	LogPath    string `mapstructure:"log_path"`
	MaxSize    int    `mapstructure:"nax_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type ConsulConfig struct {
	Host string   `mapstructure:"host"`
	Port int      `mapstructure:"port"`
	Tags []string `mapstructure:"tags"`
}

func Init() (err error) {
	//获取到配置文件的路径
	dir := GetRootStr()
	viper.SetConfigFile(dir + "/config/config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	//读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		//读取配置信息出了问题
		fmt.Printf("viper.ReadInConfig() failed:%v\n", err)
		return fmt.Errorf("viper.ReadInConfig() failed:%v\n", err)
	}
	//读取信息到global变量中
	if err = viper.Unmarshal(globalConfig); err != nil {
		fmt.Printf("viper.Unmarshal failed:%v\n", err)
		return fmt.Errorf("viper.Unmarshal failed:%v\n", err)
	}
	//实现动态配置
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改...")
		if err = viper.Unmarshal(globalConfig); err != nil {
			fmt.Printf("viper.Unmarshal failed:%v\n", err)
		}
	})
	return nil
}

func GetGlobalConfig() *GlobalConfig {
	return globalConfig
}
