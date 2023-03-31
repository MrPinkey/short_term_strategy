package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"shortTermStrategy/pkg/model"
	"strconv"
	"strings"
)

func SaveData(c *gin.Context) {
	all := findAll()
	for _, v := range all {
		data := CrawlData(v, 10000)
		if data == nil {
			continue
		}
		model.DB.CreateInBatches(&data, len(data))
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "数据插入完成",
	})
}

func findAll() []int {
	var r []int
	var s []string
	model.DB.Debug().Select("symbol").Find(&model.Lists{}).Scan(&s)
	for _, v := range s {
		a, _ := strconv.Atoi(v[2:8])
		r = append(r, a)
	}
	return r
}

var _url string = "https://d.10jqka.com.cn/v6/line/33_"

// CrawlData 获取个股历史数据
func CrawlData(code int, days int) []model.StockData {
	url := fmt.Sprintf("%s%d/01/last%d.js", _url, code, days)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//避免被禁止访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(resp)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	/*if b, _ := io.ReadAll(resp.Body); len(b) == 0 {
		resp = retry(req)
	}*/
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if strings.Contains(string(body), "Nginx forbidden") {
		return nil
	}
	hisData := parseJson(body)
	name := hisData.Name
	var ds []model.StockData
	data := hisData.Data
	split := strings.Split(data, ";")
	for i := 0; i < len(split); i++ {
		s := strings.Split(split[i], ",")
		var d = model.StockData{}
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
				if float32(cp) == d.ClosingPrice {
					d.DailyLimit = true
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
func parseJson(b []byte) model.StockHistoryData {
	d := model.StockHistoryData{}
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
func parseFloat(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}
