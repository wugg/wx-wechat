package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"wechat/config"
	"wechat/offiaccount"
)

func main() {
	var conf config.Config
	var redisConf redis.Options         // 设置配置文件类型为yaml
	viper.SetConfigFile("confing.yaml") // 设置配置文件名称
	// 添加配置文件搜索路径
	err := viper.ReadInConfig() // 读取配置文件

	if err != nil {
		panic(err)
	}
	// 检查必要的配置项是否为空
	if viper.GetString("wechat.appid") == "" || viper.GetString("wechat.appsecret") == "" {
		panic("配置文件中微信AppID或AppSecret为空")
	}
	if viper.GetString("redis.addr") == "" {
		panic("配置文件中Redis地址为空")
	}
	// 成功读取配置后赋值给全局变量
	conf.AppID = viper.GetString("wechat.appid")
	conf.AppSecret = viper.GetString("wechat.appsecret")
	redisConf.Addr = viper.GetString("redis.addr")
	redisConf.Password = viper.GetString("redis.password")
	redisConf.DB = viper.GetInt("redis.db")
	offiaccount.Instace(conf, redisConf)
	jstikcte := offiaccount.WxObj.Ticket
	fmt.Println(jstikcte)
}
