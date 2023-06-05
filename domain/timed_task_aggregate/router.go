package timed_task_aggregate

import (
	"github.com/gin-gonic/gin"
	"shortTermStrategy/domain/timed_task_aggregate/service"
)

type TimedTaskService struct {
}

func (tts *TimedTaskService) RegisterRouter(r *gin.Engine) {
	// 建表
	r.GET("create", service.CrateTables)
	// 爬取数据
	r.GET("save", service.SaveHisData)
	r.GET("saveBaseInformation", service.SaveStockInformation)
	// 从雪球获取个股历史数据
	r.GET("savaHisData", service.SaveHisDataFromXueQiu)

	// 股票概念
	r.GET("savaConceptData", service.CrawlConcepts)
}
