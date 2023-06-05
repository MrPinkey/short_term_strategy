package repository

import (
	"shortTermStrategy/domain/short_game/entity"
	infrastructure "shortTermStrategy/infrastructure/config"
)

func GetDailyLimit() []entity.StockData {
	var results []entity.StockData
	var sh []entity.StockData
	var sz []entity.StockData
	infrastructure.DB.Table("stock_his_data_sh").
		Where("stock_his_data_sh.daily_limit = ? AND stock_his_data_sh.date > ? AND stock_his_data_sh.date < ?", "1", "2023-05-01", "2023-05-20").
		Select("stock_his_data_sh.date, stock_his_data_sh.stock_name, CONCAT('SH', stock_his_data_sh.stock_code) AS stock_code").
		Scan(&sh)
	infrastructure.DB.Table("stock_his_data_sz").
		Where("stock_his_data_sz.daily_limit = ? AND stock_his_data_sz.date > ? AND stock_his_data_sz.date < ?", "1", "2023-05-01", "2023-05-20").
		Select("stock_his_data_sz.date, stock_his_data_sz.stock_name, stock_his_data_sz.stock_code").
		Scan(&sz)
	results = append(sh, sz...)
	return results
}

func GetConcepts(str string) []entity.StockConcept {
	var results []entity.StockConcept
	infrastructure.DB.Table("stock_concepts").
		Where("symbol = ?", str).
		Select("concept, region, plate").
		Scan(&results)
	return results
}
