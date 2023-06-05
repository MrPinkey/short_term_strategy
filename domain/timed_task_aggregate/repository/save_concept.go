package repository

import (
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	infrastructure "shortTermStrategy/infrastructure/config"
)

func SaveStockConcept(c string, r string, p string, cp []string) {
	for _, v := range cp {
		infrastructure.DB.Create(&entity.StockConcept{
			Symbol:  c,
			Region:  r,
			Plate:   p,
			Concept: v,
		})
	}
}

func GetAllSymbol() []string {
	var r []string
	infrastructure.DB.Select("symbol").
		Find(&entity.BaseInfo{}).Scan(&r)
	return r
}
