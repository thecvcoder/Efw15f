package main

import (
	"github.com/thecvcoder/cloud-api-go/core"
	"github.com/thecvcoder/cloud-api-go/global"
	"github.com/thecvcoder/cloud-api-go/initialize"
)

func init() {
	// 初始化Viper
	global.Viper = initialize.InitViper()
	//  初始化日志
	global.LOG = initialize.InitLogger()
	//  初始化mysql数据库
	global.DB = initialize.InitMysql()
	//  初始化数据
	initialize.InitData()
	// 初始化Validator
	initialize.InitValidator()
	//初始化redis
	//initialize.InitRedis()
}

func main() {
	//  启动服务
	core.RunServer()
}
