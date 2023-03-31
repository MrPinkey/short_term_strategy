package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"shortTermStrategy/pkg/model"
)

func SaveStockInformation(c *gin.Context) {
	si := getStockInformation()
	model.DB.CreateInBatches(&si, len(si))
}

func getStockInformation() []model.Lists {
	var stock model.Stock
	url := "https://xueqiu.com/service/screener/screen?category=CN&exchange=sha&size=1679&only_count=0&page=1"
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
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &stock)
	if err != nil {
		log.Fatalln("股票基本数据解析出错")
	}
	return stock.AllData.List
}
