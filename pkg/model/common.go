package model

import "time"

type StockHistoryData struct {
	Total string `json:"total"`
	Start string `json:"start"`
	Year  Years  `json:"year"`
	Name  string `json:"name"`
	Data  string `json:"data"`
}

type Years struct {
	Year string
	days int
}

type StockData struct {
	ID                 int32   `gorm:"column:id;not null;primary key;auto_increment"`
	StockCode          string  `gorm:"column:stock_code;not null;comment:股票代码"`
	StockName          string  `gorm:"column:stock_name;comment:股票名称"`
	Date               string  `gorm:"column:date;comment:股票日期"`
	OpeningPrice       float32 `gorm:"opening_price;comment:开盘价"`
	HighestPrice       float32 `gorm:"highest_price;comment:最高价"`
	LowestPrice        float32 `gorm:"lowest_price;comment:最低价"`
	ClosingPrice       float32 `gorm:"closing_price;comment:收盘价"`
	NumberTransactions int     `gorm:"number_transactions;comment:成交量"`
	TurnoverRate       float32 `gorm:"turnover_rate;comment:换手率"`
	DailyLimit         bool    `gorm:"daily_limit;comment:涨停标志"`
	CreatedAt          time.Time
}

func (sd *StockData) TableName() string {
	return "stock_data_sh"
}

type Stock struct {
	AllData     Data   `json:"data"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"error_description"`
}

type Data struct {
	Count int `json:"count"`
	List  []Lists
}

type Lists struct {
	ID        int32   `gorm:"column:id;not null;primary key;auto_increment"`
	Pct       float32 `json:"pct" gorm:"pct;comment:股票名称"`
	Symbol    string  `json:"symbol" gorm:"symbol;comment:股票代码"`
	Name      string  `json:"name" gorm:"name;comment:股票名称"`
	CreatedAt time.Time
}

func (sd *Lists) TableName() string {
	return "stock_information_sh"
}
