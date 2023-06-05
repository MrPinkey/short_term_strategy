package repository

import (
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	"shortTermStrategy/infrastructure/config"
	"strconv"
	"strings"
)

// SaveData 插入股票历史数据表
func SaveData(s string, data []entity.HisData) {
	if s == "sha" {
		infrastructure.DB.Table("stock_his_data_sh").CreateInBatches(&data, len(data))
	} else if s == "sza" {
		infrastructure.DB.Table("stock_his_data_sz").CreateInBatches(&data, len(data))
	}
}

// FindAll 获取所有股票代码
func FindAll(s string) []int {
	var i []int
	var r []string
	var stockCodes []string
	result := infrastructure.DB.Table("stock_his_data_" + s[:2]).Distinct("stock_code")
	if result.Error != nil {
		panic(result.Error)
	}
	result.Pluck("stock_code", &stockCodes)
	for j := 0; j < len(stockCodes); j++ {
		stockCodes[j] = strings.ToUpper(s[:2]) + stockCodes[j]
	}
	/*infrastructure.DB.Select("symbol").
	Where("area = ?", s).
	Not(infrastructure.DB.Where("name IN (?)", stockNames)).
	Find(&entity.BaseInfo{}).Scan(&r)*/
	if len(stockCodes) > 0 {
		infrastructure.DB.Select("symbol").
			Where("area = ?", s).
			Not(infrastructure.DB.Where("symbol IN (?)", stockCodes)).
			Find(&entity.BaseInfo{}).Scan(&r)
	} else {
		infrastructure.DB.Select("symbol").
			Where("area = ?", s).
			Find(&entity.BaseInfo{}).Scan(&r)
	}
	for _, v := range r {
		a, _ := strconv.Atoi(v[2:8])
		i = append(i, a)
	}
	return i
}
