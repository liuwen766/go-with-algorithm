package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	fmt.Println("Viper是适用于Go应用程序的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式。")
	fmt.Println("Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目:\n显示调用Set设置值\n命令行参数（flag）\n环境变量\n配置文件\nkey/value存储\n默认值\n")

	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	viper.SetConfigFile("./config.yaml")  // 指定配置文件路径
	viper.SetConfigName("config")         // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")           // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("/etc/appName/")  // 查找配置文件所在的路径
	viper.AddConfigPath("$HOME/.appName") // 多次调用以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()           // 查找并读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
		} else {
			// 配置文件被找到，但产生了另外的错误
		}
	}
	// 配置文件找到并成功解析

	viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
	viper.SafeWriteConfig()
	viper.WriteConfigAs("/path/to/my/.config")
	viper.SafeWriteConfigAs("/path/to/my/.config") // 因为该配置文件写入过，所以会报错
	viper.SafeWriteConfigAs("/path/to/my/.other_config")

	fmt.Println("Viper 有五种方法可以帮助与ENV环境变量协作:\nAutomaticEnv()\nBindEnv(string...) : error\nSetEnvPrefix(string)\nSetEnvKeyReplacer(string...) *strings.Replacer\nAllowEmptyEnv(bool)")

	fmt.Println("Viper 具有绑定到标志flag的能力。如：Pflag")
	pflag.Int("flagname", 1234, "help message for flagname")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	i := viper.GetInt("flagname") // 从viper而不是从pflag检索值
	fmt.Println(i)

	fmt.Println("Viper 远程Key/Value存储示例")
	fmt.Println("需要 Consul Key/Value存储中设置一个Key保存包含所需配置的JSON值。如：\n{\n\t\t\"port\": 8080,\n\t\t\"hostname\": \"liwenzhou.com\"\n}")
	viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
	viper.SetConfigType("json") // 需要显示设置成json
	err = viper.ReadRemoteConfig()
	fmt.Println(viper.Get("port"))     // 8080
	fmt.Println(viper.Get("hostname")) // liwenzhou.com
}
