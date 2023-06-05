package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shortTermStrategy/domain/auth_aggregate/entity"
)

func GetMenus(c *gin.Context) {
	var result entity.Result
	var menus = make([]entity.Menu, 0)
	menus = append(menus, entity.Menu{
		Id:       1,
		AuthName: "短期博弈",
		Path:     "shortGame",
		Children: []entity.Menu{
			{
				Id:       11,
				AuthName: "近期涨停仅1次",
				Path:     "once",
				Children: []entity.Menu{},
				Order:    1,
			},
			{
				Id:       12,
				AuthName: "近期涨停超1次",
				Path:     "more",
				Children: []entity.Menu{},
				Order:    2,
			},
		},
		Order: 1,
	})
	menus = append(menus, entity.Menu{
		Id:       2,
		AuthName: "角色权限",
		Path:     "permissions",
		Children: []entity.Menu{
			{
				Id:       21,
				AuthName: "角色列表",
				Path:     "roles",
				Children: []entity.Menu{},
				Order:    1,
			},
			{
				Id:       22,
				AuthName: "权限列表",
				Path:     "rights",
				Children: []entity.Menu{},
				Order:    2,
			},
		},
		Order: 2,
	})
	menus = append(menus, entity.Menu{
		Id:       3,
		AuthName: "会员服务",
		Path:     "members",
		Children: []entity.Menu{
			{
				Id:       31,
				AuthName: "许愿池",
				Path:     "wishing",
				Children: []entity.Menu{},
				Order:    1,
			},
			{
				Id:       32,
				AuthName: "还愿池",
				Path:     "votive",
				Children: []entity.Menu{},
				Order:    2,
			},
		},
		Order: 3,
	})
	menus = append(menus, entity.Menu{
		Id:       4,
		AuthName: "资讯中心",
		Path:     "news",
		Children: []entity.Menu{
			{
				Id:       41,
				AuthName: "国内新闻",
				Path:     "domestic",
				Children: []entity.Menu{},
				Order:    1,
			},
			{
				Id:       42,
				AuthName: "国外新闻",
				Path:     "foreign",
				Children: []entity.Menu{},
				Order:    2,
			},
		},
		Order: 4,
	})
	menus = append(menus, entity.Menu{
		Id:       5,
		AuthName: "交易策略",
		Path:     "strategy",
		Children: []entity.Menu{
			{
				Id:       51,
				AuthName: "左侧",
				Path:     "left",
				Children: []entity.Menu{},
				Order:    1,
			},
			{
				Id:       52,
				AuthName: "右侧",
				Path:     "right",
				Children: []entity.Menu{},
				Order:    2,
			},
		},
		Order: 5,
	})
	result.Data = menus
	result.Meta.Status = 200
	result.Meta.Message = "获取菜单列表成功"
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
