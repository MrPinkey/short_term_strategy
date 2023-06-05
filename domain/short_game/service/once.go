package service

import (
	"github.com/gin-gonic/gin"
	"shortTermStrategy/domain/short_game/entity"
	"shortTermStrategy/domain/short_game/repository"
)

func GetOnce(c *gin.Context) {
	stockData := repository.GetDailyLimit()
	// 使用 map 将具有相同 StockName 值的元素分组
	groups := make(map[string][]entity.StockData)
	for _, s := range stockData {
		if _, ok := groups[s.StockName]; ok {
			// 如果已经存在这个 StockName 的分组，则将当前元素放入这个分组的切片中
			groups[s.StockName] = append(groups[s.StockName], s)
		} else {
			// 如果不存在这个 StockName 的分组，则创建一个新的切片，并将当前元素放入其中
			groups[s.StockName] = []entity.StockData{s}
		}
	}
	// 遍历 map 中的每个分组，将分组中元素个数大于 1 的放入一个切片中，其余的放入另一个切片中
	var singles []entity.StockData
	for _, group := range groups {
		if len(group) == 1 {
			// 如果分组中只有一个元素，则将其放入 singles 切片中
			singles = append(singles, group[0])
		}
	}
	var res []entity.DailyLimit
	for _, v := range singles {
		var dl entity.DailyLimit
		concepts := repository.GetConcepts(v.StockCode)
		for _, concept := range concepts {
			dl.Concepts = append(dl.Concepts, concept.Concept)
			dl.Region = concept.Region
			dl.Plate = concept.Plate
			dl.StockName = v.StockName
			dl.Date = v.Date
		}
		res = append(res, dl)
	}
	var result entity.Result
	result.Data = res
	result.Meta.Status = 200
	result.Meta.Message = "success"
	c.JSON(200, gin.H{
		"result": result,
	})
}
