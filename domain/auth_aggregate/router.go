package auth_aggregate

import (
	"github.com/gin-gonic/gin"
	"shortTermStrategy/domain/auth_aggregate/service"
)

type AuthService struct {
}

func (as *AuthService) RegisterRouter(r *gin.Engine) {
	//登录页面验证
	r.POST("auth", service.VerifyPermissions)
	//获取菜单列表
	r.GET("getMenus", service.GetMenus)
}
