package repository

import (
	"log"
	infrastructure "shortTermStrategy/infrastructure/config"
)

func GetCodeName(s string) string {
	var name *string
	tx := infrastructure.DB.Table("stock_base_information").Where("symbol = ?", s).Select("name").Scan(&name)
	log.Println(tx.Error)
	return *name
}
