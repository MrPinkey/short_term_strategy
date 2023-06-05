package service

import (
	"github.com/gin-gonic/gin"
	"shortTermStrategy/domain/timed_task_aggregate/repository"
)

func CrateTables(c *gin.Context) {
	err := repository.CreateTable()
	if err == nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "创建成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "创建失败",
		})
	}
}
