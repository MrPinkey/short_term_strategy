package infrastructure

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
)

var DB *gorm.DB

func init() {
	var base DBConfig
	c, _ := ioutil.ReadFile("config.yaml")
	_ = yaml.Unmarshal(c, &base)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		base.MySQL.User, base.MySQL.Password, base.MySQL.Host, base.MySQL.Port, base.MySQL.Name)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db, err := _db.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	if err != nil {
		log.Fatalln("mysql connected error: ", err)
	}
	DB = _db
}
