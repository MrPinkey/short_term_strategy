package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"shortTermStrategy/domain/timed_task_aggregate/entity"
	"shortTermStrategy/domain/timed_task_aggregate/repository"
	"strings"
)

func CrawlConcepts(ct *gin.Context) {
	symbol := repository.GetAllSymbol()
	var data entity.ConceptData
	for _, item := range symbol {
		_con := "https://emweb.securities.eastmoney.com/PC_HSF10/CoreConception/PageAjax?code="
		url := fmt.Sprintf("%s%s#", _con, item)
		// 发送 HTTP 请求获取网页内容
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		all, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(all, &data)
		if err != nil {
			panic(err)
		}
		var concept []string
		plate := data.Ssbk[0].BoardName
		region := data.Ssbk[1].BoardName
		for _, v := range data.Ssbk[2:] {
			concept = append(concept, v.BoardName)
		}
		repository.SaveStockConcept(item, region, plate, concept)
	}
	ct.JSON(http.StatusOK, gin.H{
		"message": "股票概念数据插入完成",
	})
}

// CrawlConceptsBackup 下面的网页可能有所属地为空的情况，暂时弃用
func CrawlConceptsBackup(ct *gin.Context) {
	symbol := repository.GetAllSymbol()
	_con := "https://stockpage.10jqka.com.cn/"
	for _, item := range symbol {
		url := fmt.Sprintf("%s%s", _con, item[2:8])
		// 发送 HTTP 请求获取网页内容
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var s string
		var cs []string
		// 从 HTML 文档中抽取概念
		var findTitle func(*html.Node, *bool)
		findTitle = func(n *html.Node, found *bool) {
			if n.Type == html.ElementNode && n.Data == "dd" {
				if !*found {
					for _, attr := range n.Attr {
						if attr.Key == "title" {
							cs = strings.Split(attr.Val, "，")
							//fmt.Println(attr.Val)
							*found = true
							return
						}
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				findTitle(c, found)
			}
		}
		// 从 HTML 文档中抽取所属地区
		var getDdValue func(*html.Node) string
		getDdValue = func(n *html.Node) string {
			if n.Type == html.ElementNode && n.Data == "dd" {
				return strings.TrimSpace(n.FirstChild.Data)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if v := getDdValue(c); v != "" {
					return v
				}
			}
			return ""
		}
		found := false
		findTitle(doc, &found)
		value := getDdValue(doc)
		s = value
		// 保存到数据库
		repository.SaveStockConcept(item, s, "", cs)

		func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(resp.Body)
	}
	ct.JSON(http.StatusOK, gin.H{
		"message": "概念保存成功",
	})
}
