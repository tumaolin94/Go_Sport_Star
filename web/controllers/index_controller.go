// file: controllers/index_controller.go

package controllers

import (
	"github.com/kataras/iris/mvc"
	"log"
	"superstar/datasource"
	"superstar/models"
	"superstar/services"

	"github.com/kataras/iris"
)

// UserController is our /admin controller.
// UserController is responsible to handle the following requests:
// GET  			/admin/register
// POST 			/admin/register
// GET 				/admin/login
// POST 			/admin/login
// GET 				/admin/me
// All HTTP Methods /admin/logout
type IndexController struct {
	Ctx iris.Context

	Service services.SuperstarService

}

func (c *IndexController) Get() mvc.Result{
	//log.Fatal("Get22222")
	datalist := c.Service.GetAll()
	return mvc.View{
		Name:"index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Datalist": datalist,
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

func (c *IndexController) GetSearch() mvc.Result{
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	return mvc.View{
		Name:"index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Datalist": datalist,
		},
	}
}
// http://localhost:8080/clearcache
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatal(err)
	}
	// set the model and render the view template.
	return mvc.Response{
		Text: "xorm clear catch",
	}
}