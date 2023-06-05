package entity

type StockConcept struct {
	Concept string `gorm:"column:concept"`
	Region  string `gorm:"column:region;not null;default:''"`
	Plate   string `gorm:"column:plate;not null;default:''"`
}

type StockData struct {
	StockCode string `json:"stock_code"`
	StockName string `json:"stock_name"`
	Date      string `json:"date"`
}

type DailyLimit struct {
	StockName string   `json:"stock_name"`
	Date      string   `json:"date"`
	Callback  string   `json:"callback"`
	Concepts  []string `json:"concepts"`
	Region    string   `json:"region"`
	Plate     string   `json:"plate"`
}
type Result struct {
	Data []DailyLimit `json:"data"`
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
}
