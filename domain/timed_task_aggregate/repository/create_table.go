package repository

import (
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	"shortTermStrategy/infrastructure/config"
)

func CreateTable() error {
	//创建历史数据表，两种方式
	err := infrastructure.DB.Table("stock_his_data_sh").AutoMigrate(new(entity.HisData))
	err = infrastructure.DB.Table("stock_his_data_sz").AutoMigrate(&entity.HisData{})
	//创建基本信息表
	err = infrastructure.DB.AutoMigrate(&entity.BaseInfo{})
	//创建概念表
	err = infrastructure.DB.Table("stock_concepts").AutoMigrate(&entity.StockConcept{})
	return err
}
