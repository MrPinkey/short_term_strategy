package controller

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	//建表
	r.GET("create", CreateTable)
	//爬取数据
	r.GET("save", SaveData)
	r.GET("saveBaseInformation", SaveStockInformation)
}
