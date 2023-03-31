package router

import (
	"github.com/gin-gonic/gin"
	"shortTermStrategy/pkg/controller"
)

func InitRouter(r *gin.Engine) {
	controller.RegisterRouter(r)
}
