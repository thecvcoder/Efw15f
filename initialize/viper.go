package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/thecvcoder/cloud-api-go/global"
	"os"
)

func InitViper() *viper.Viper {
	workDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取应用根目录失败: %v", err))
	}

	v := viper.New()
	v.SetConfigName("config")      // 配置文件名
	v.SetConfigType("yml")         // 配置文件类型
	v.AddConfigPath(workDir + "/") // 配置文件所在目录
	if err = v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败:%s", err))
	}
	// 配置文件热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生了改变: ", e.Name)
		if err = viper.Unmarshal(&global.CONFIG); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}
	})

	if err = v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("初始化配置文件失败: %v", err))
	}

	return v
}
