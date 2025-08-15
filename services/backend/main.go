package main

import (
	"backend/config"
	"backend/dal"
	"backend/router"
	"backend/utils"
)


func main() {
	config.InitConfig() // 配置初始化

	utils.Init() // 工具初始化
	dal.Init() // 数据访问初始化

	// h := router.InitRouter()
	// h.Spin()
	router.InitRouter().Spin() 
}
