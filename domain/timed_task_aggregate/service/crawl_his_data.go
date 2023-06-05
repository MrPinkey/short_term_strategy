package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	"shortTermStrategy/domain/timed_task_aggregate/repository"
	"strconv"
	"strings"
)

func SaveHisData(c *gin.Context) {
	sha := repository.FindAll("sha")
	for _, v := range sha {
		data := CrawlData(v, 1000)
		repository.SaveData("sha", data)
	}
	sza := repository.FindAll("sza")
	for _, v := range sza {
		data := CrawlData(v, 1000)
		repository.SaveData("sza", data)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "历史数据插入成功",
	})
}

var _url = "https://d.10jqka.com.cn/v6/line/33_"

// CrawlData 获取个股历史数据
func CrawlData(code int, days int) []entity.HisData {
	url := fmt.Sprintf("%s%d/01/last%d.js", _url, code, days)
	body := GetInfoFromUrl(url)
	if strings.Contains(string(body), "Nginx forbidden") {
		return nil
	}
	hisData := parseJson(body)
	name := hisData.Name
	var ds []entity.HisData
	data := hisData.Data
	split := strings.Split(data, ";")
	for i := 0; i < len(split); i++ {
		s := strings.Split(split[i], ",")
		var d = entity.HisData{}
		for j := 0; j < len(s); j++ {
			d.Date = s[0]
			d.StockName = name
			d.StockCode = strconv.Itoa(code)
			d.OpeningPrice = parseFloat(s[1])
			d.HighestPrice = parseFloat(s[2])
			d.LowestPrice = parseFloat(s[3])
			d.ClosingPrice = parseFloat(s[4])
			d.NumberTransactions, _ = strconv.Atoi(s[5])
			d.TurnoverRate = parseFloat(s[7])
			//涨停字段值设置
			if i > 1 {
				sd := ds[i-1]
				cp, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", sd.ClosingPrice*1.1), 64)
				if cp == d.ClosingPrice {
					d.DailyLimit = 1
				}
			}
		}
		ds = append(ds, d)
	}
	return ds
}

// 重试
func retry(r *http.Request) *http.Response {
	client := &http.Client{}
	resp, _ := client.Do(r)
	return resp
}

// 解析字符串
func parseJson(b []byte) entity.StockHistoryData {
	d := entity.StockHistoryData{}
	i := strings.Index(string(b), "(")
	l := len(string(b))
	r := string(b)[i+1 : l-1]
	err := json.Unmarshal([]byte(r), &d)
	if err != nil {
		log.Fatalln(err)
	}
	return d
}

// 类型转换
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 32)
	return f
}
