package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	"shortTermStrategy/domain/timed_task_aggregate/repository"
)

func SaveStockInformation(c *gin.Context) {
	saveInformation("sha")
	saveInformation("sza")
	c.JSON(http.StatusOK, gin.H{
		"result": "基础数据插入成功",
	})
}

func saveInformation(s string) {
	cs := getCounts(s)
	n := cs/200 + 1
	for i := 1; i <= n; i++ {
		var stock entity.Stock
		url := fmt.Sprintf("https://xueqiu.com/service/screener/screen?category=CN&exchange=%s&only_count=0&page=%d&size=%d", s, i, 200)
		body := GetInfoFromUrl(url)
		err := json.Unmarshal(body, &stock)
		if err != nil {
			log.Fatalln("股票基本数据解析出错")
		}
		for j := range stock.Data.List {
			stock.Data.List[j].Area = s
		}
		repository.SaveStockInformation(stock)
	}
}

// GetCounts 获取股票总数
func getCounts(s string) int {
	type Response struct {
		Data struct {
			Count int `json:"count"`
		} `json:"data"`
		ErrorCode        int    `json:"error_code"`
		ErrorDescription string `json:"error_description"`
	}
	url := "https://xueqiu.com/service/screener/screen?category=CN&exchange=" + s
	body := GetInfoFromUrl(url)
	var res Response
	err := json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalln("请求出错")
	}
	fmt.Println("股票总数：", res.Data.Count)
	return res.Data.Count
}

// GetInfoFromUrl 从url获取数据
func GetInfoFromUrl(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// 访问官网先获取cookie，再添加cookie到请求头
	res := getCookies()
	cookies := res.Cookies()
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(resp)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}

func getCookies() *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://xueqiu.com", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
	return resp
}
