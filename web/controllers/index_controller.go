package controllers

import (
	"superstar/services"

	"github.com/kataras/iris/v12/mvc"

	"github.com/kataras/iris/v12"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.SuperStarService
}

func (c *IndexController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"DataList": dataList,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"info":  data,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	dataList := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"DataList": dataList,
		},
	}
}
