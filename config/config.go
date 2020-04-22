package config

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"path"
	"strings"
)

//初始化配置，通过配置文件载入各个配置变量
func Initialize() error {
	viper.Reset()
	viper.SetEnvPrefix("zues")
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName("zues")
	viper.SetConfigType("yaml")
	homepath := viper.GetString("rule")
	if homepath != "" {
		viper.AddConfigPath(homepath)
	} else {
		return fmt.Errorf("Not find zues yaml")
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}

//GetString以字符串形式返回与键关联的值
func GetString(key string) string {
	return viper.GetString(key)
}

//GetInt以Int形式返回与键关联的值
func GetInt(key string) int {
	return viper.GetInt(key)
}

//GetInt64以Int64形式返回与键关联的值
//灰色表示没有被调用，全局方法没有被调用不会报错
func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

//GetInt32以int32形式返回与键关联的值
func GetInt32(key string) int32 {
	return cast.ToInt32(viper.GetInt(key))
}

//拼接与键关联的路径
func GetPath(key string) string {
	return path.Join(viper.GetString("home"), viper.GetString(key))
}

//GetMap以Map形式返回与键关联的值
func GetMap(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

//GetStringSlice方法将与键关联的接口类型的值强制转换为[]string类型
func GetStringSlice(key string) []string {
	return cast.ToStringSlice(viper.GetString(key))
}

//GetBool以布尔形式返回与键关联的值
func GetBool(key string) bool {
	return viper.GetBool(key)
}

//Get方法可以检索给定键对应的任何值
func Get(key string) interface{} {
	return viper.Get(key)
}

//Set方法设置键的值
func Set(key string, value interface{}) {
	viper.Set(key, value)
}

//GetStringMap将与键关联的接口强制转换为map[string]interface{}类型
func GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(Get(key))
}
