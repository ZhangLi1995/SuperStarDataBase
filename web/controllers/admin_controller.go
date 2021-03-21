package controllers

import (
	"log"
	"superstar/models"
	"superstar/services"
	"time"

	"github.com/kataras/iris/v12/mvc"

	"github.com/kataras/iris/v12"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperStarService
}

func (c *AdminController) Get() mvc.Result {
	dataList := c.Service.GetAll()
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"DataList": dataList,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	var data *models.StarInfo
	if id, err := c.Ctx.URLParamInt("id"); err == nil {
		data = c.Service.Get(id)
	}
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html",
	}
}

func (c *AdminController) PostSave() mvc.Result {
	star := &models.StarInfo{}
	err := c.Ctx.ReadForm(star)
	if err != nil {
		log.Fatal(err)
	}
	if star.Id > 0 {
		star.SysUpdated = int(time.Now().Unix())
		c.Service.Update(star, []string{"name_zh", "name_en", "avatar", "birthday", "height", "weight", "club",
			"jersy", "country", "moreinfo"})
	} else {
		star.SysCreated = int(time.Now().Unix())
		c.Service.Create(star)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	if id, err := c.Ctx.URLParamInt("id"); err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}
