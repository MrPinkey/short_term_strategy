package entity

import (
	"time"
)

type StockHistoryData struct {
	Total string `json:"total"`
	Start string `json:"start"`
	Years struct {
		Year string
		Days int
	} `json:"year"`
	Name string `json:"name"`
	Data string `json:"data"`
}

type HisData struct {
	ID                 int64   `gorm:"column:id;not null;primary key;auto_increment"`
	StockCode          string  `gorm:"column:stock_code;not null;comment:股票代码;index:idx_stock_code"`
	StockName          string  `gorm:"column:stock_name;comment:股票名称;index:idx_stock_name"`
	Date               string  `gorm:"column:date;comment:股票日期;index:idx_date"`
	OpeningPrice       float64 `gorm:"opening_price;comment:开盘价"`
	HighestPrice       float64 `gorm:"highest_price;comment:最高价"`
	LowestPrice        float64 `gorm:"lowest_price;comment:最低价"`
	ClosingPrice       float64 `gorm:"closing_price;comment:收盘价"`
	NumberTransactions int     `gorm:"number_transactions;comment:成交量"`
	TurnoverRate       float64 `gorm:"turnover_rate;comment:换手率"`
	DailyLimit         int     `gorm:"daily_limit;comment:涨停标志;index:idx_daily_limit"`
	CreatedAt          time.Time
}

type Stock struct {
	Data struct {
		Count int `json:"count"`
		List  []BaseInfo
	} `json:"data"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"error_description"`
}

type BaseInfo struct {
	ID        int32  `gorm:"column:id;not null;primary key;auto_increment"`
	Symbol    string `json:"symbol" gorm:"symbol;comment:股票代码"`
	Name      string `json:"name" gorm:"name;comment:股票名称"`
	Area      string `json:"area" gorm:"area;comment:地区"`
	CreatedAt time.Time
}

func (sd *BaseInfo) TableName() string {
	return "stock_base_information"
}

type ConceptTable struct {
	ID        int32  `gorm:"column:id;not null;primary key;auto_increment"`
	StockCode string `json:"stock_code" gorm:"stock_code;comment:股票代码"`
	StockName string `json:"stock_name" gorm:"stock_name;comment:股票名称"`
	Industry  string `json:"industry" gorm:"industry;comment:行业"`
	Region    string `json:"region" gorm:"region;comment:地区"`
	Concept   string `json:"concept" gorm:"concept;comment:概念"`
}

type StockConcept struct {
	ID        int32  `gorm:"column:id;not null;primary key;auto_increment"`
	Symbol    string `json:"symbol" gorm:"symbol;comment:股票代码"`
	Concept   string `json:"concept" gorm:"concept;comment:概念"`
	Region    string `json:"region" gorm:"region;comment:地区"`
	Plate     string `json:"plate" gorm:"plate;comment:板块"`
	CreatedAt time.Time
}
