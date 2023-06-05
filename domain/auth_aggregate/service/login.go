package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortTermStrategy/domain/auth_aggregate/entity"
)

func VerifyPermissions(c *gin.Context) {
	var user entity.User
	var res entity.Res
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if user.Username == "admin" && user.Password == "123456" {
		res.Data.Name = "admin"
		res.Data.Age = 18
		res.Data.Email = "XX@qq.com"
		res.Data.Token = "token"
		res.Meta.Status = 200
		res.Meta.Message = "登录成功"
		c.JSON(http.StatusOK, gin.H{
			//使用res作为返回结果，随便给出一些值
			"result": res,
		})
	} else {
		res.Meta.Status = 400
		res.Meta.Message = "登录失败"
		c.JSON(http.StatusOK, gin.H{
			"result": res,
		})
	}
}
