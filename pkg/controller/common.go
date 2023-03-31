package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortTermStrategy/pkg/model"
)

func CreateTable(c *gin.Context) {
	//创建基础数据表和历史数据表
	model.DB.AutoMigrate(new(model.Lists), new(model.StockData)).Error()
	model.DB.AutoMigrate().Error()
	c.JSON(http.StatusOK, gin.H{
		"msg": "建表成功",
	})
}
