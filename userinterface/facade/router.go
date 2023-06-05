package userinterface

import (
	"github.com/gin-gonic/gin"
	application "shortTermStrategy/application/service"
)

func ServiceGroupApp(r *gin.Engine) {
	serviceApplication := application.ServiceApplication{}
	serviceApplication.TimedTask.RegisterRouter(r)
	serviceApplication.AuthService.RegisterRouter(r)
	serviceApplication.ShortService.RegisterRouter(r)
}
