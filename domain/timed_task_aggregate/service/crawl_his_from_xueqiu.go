package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	"shortTermStrategy/domain/timed_task_aggregate/repository"
	"strconv"
	"strings"
	"time"
)

func SaveHisDataFromXueQiu(c *gin.Context) {
	sha := repository.FindAll("sha")
	for _, v := range sha {

		data := ObtainHistoricalDataFromXueQiu(v, -12000, "SH")
		var batchSize = 5000                                     // 每个切片的最大长度
		var numBatches = (len(data) + batchSize - 1) / batchSize // 计算切分后的切片数量
		for i := 0; i < numBatches; i++ {
			start := i * batchSize     // 计算当前切片的起始索引
			end := (i + 1) * batchSize // 计算当前切片的结束索引
			if end > len(data) {       // 如果结束索引超出了切片的长度，就将其设置为切片的最后一个元素的下一个位置
				end = len(data)
			}
			batch := data[start:end]
			repository.SaveData("sha", batch)
		}
	}
	sza := repository.FindAll("sza")
	for _, v := range sza {
		data := ObtainHistoricalDataFromXueQiu(v, -12000, "SZ")
		var batchSize = 5000
		var numBatches = (len(data) + batchSize - 1) / batchSize
		for i := 0; i < numBatches; i++ {
			start := i * batchSize
			end := (i + 1) * batchSize
			if end > len(data) {
				end = len(data)
			}
			batch := data[start:end]
			repository.SaveData("sza", batch)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "历史数据插入成功",
	})
}

type Data struct {
	Symbol string      `json:"symbol"`
	Column []string    `json:"column"`
	Item   [][]float64 `json:"item"`
}

type Response struct {
	Data             Data   `json:"data"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

var _xueQiu = "https://stock.xueqiu.com/v5/stock/chart/kline.json?period=day&type=before&indicator=kline,pe,pb,ps,pcf,market_capital,agt,ggt,balance&symbol="

// ObtainHistoricalDataFromXueQiu 获取个股历史数据
func ObtainHistoricalDataFromXueQiu(code int, days int, area string) []entity.HisData {
	// 获取当前本地时间戳，单位为秒
	timestamp := time.Now().Unix()
	// 将时间戳乘以 1000，得到以毫秒为单位的时间戳
	timestampMillis := timestamp * 1000
	// 将时间戳格式化为字符串
	timestampStr := fmt.Sprintf("%d", timestampMillis)
	// 爬取个股历史数据
	sc := strconv.Itoa(code)
	symbol := area + fmt.Sprintf("%06s", sc)
	url := fmt.Sprintf("%s%s&begin=%s&count=%d", _xueQiu, symbol, timestampStr, days)
	body := GetInfoFromUrl(url)
	if strings.Contains(string(body), "Nginx forbidden") || strings.Contains(string(body), "遇到错误") {
		return nil
	}
	var resp Response
	err := json.Unmarshal(body, &resp)
	if err != nil {
		panic(err)
	}
	codeName := repository.GetCodeName(symbol)
	var ds []entity.HisData
	// 遍历 item 数组中的所有元素
	for i, row := range resp.Data.Item {
		var d = entity.HisData{}
		d.StockCode = symbol
		d.StockName = codeName
		// 将时间戳转换为日期
		tm := time.Unix(int64(row[0]/1000), 0)
		d.Date = tm.Format("2006-01-02")
		d.OpeningPrice = math.Round(row[2]*100) / 100
		d.HighestPrice = math.Round(row[3]*100) / 100
		d.LowestPrice = math.Round(row[4]*100) / 100
		d.ClosingPrice = math.Round(row[5]*100) / 100
		d.NumberTransactions = int(row[9])
		d.TurnoverRate = row[8]
		if i > 4 {
			sd := ds[i-1]
			// 涨停价格
			dailyLimit, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", sd.ClosingPrice*1.1), 64)
			// 跌停价格
			downLimit, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", sd.ClosingPrice*0.9), 64)
			if dailyLimit == d.ClosingPrice {
				d.DailyLimit = 1
			} else if downLimit == d.ClosingPrice {
				d.DailyLimit = -1
			} else {
				d.DailyLimit = 0
			}
		}
		t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
		//bj := t.In(time.FixedZone("CST", 8*3600))
		d.CreatedAt = t
		ds = append(ds, d)
	}
	return ds
}
