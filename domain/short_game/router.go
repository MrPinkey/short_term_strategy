package shortGame

import (
	"github.com/gin-gonic/gin"
	"shortTermStrategy/domain/short_game/service"
)

type ShortService struct {
}

func (ss *ShortService) RegisterRouter(r *gin.Engine) {
	r.GET("getOnce", service.GetOnce)
}
