package cmd


import (
	"fmt"
	"strings"

	"github.com/spf13/viper"

)


var Config GmctlConfig

type GmctlConfig struct {
	Server  string
	Token string
}

//func init() {
//	ConfigSetup()
//}


//https://www.jianshu.com/p/74c933f80c6f
//https://copyfuture.com/blogs-details/202204130241322101
func ConfigSetup() {
	// 注入配置扩展项
	//var Config *viper.Viper
	//fmt.Println("Using config file:", ConfigFile)
	//Config = viper.New()
	viper.SetConfigFile(ConfigFile)
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//解析到结构体中

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("err:",err )
	}

	fmt.Println("allkeys: ",Config)

	// For environment variables.
	viper.AutomaticEnv()  //是否去读取环境变量
	viper.SetEnvPrefix("TEST")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	key := "A1.B"
	fmt.Printf("v[%s]=[%s]\n", key, viper.GetString(key))
	key  = "A2.B"
	fmt.Printf("v[%s]=[%s]\n", key, viper.GetString(key))


}
