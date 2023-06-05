package repository

import (
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	infrastructure "shortTermStrategy/infrastructure/config"
)

func SaveStockInformation(d entity.Stock) {
	infrastructure.DB.CreateInBatches(&d.Data.List, len(d.Data.List))
}
