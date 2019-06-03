// file: controllers/index_controller.go

package controllers

import (
	"github.com/kataras/iris/mvc"
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
	datalist := c.Service.GetAll()
	return mvc.View{
		Name:"index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result{
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name:"index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Data": data,
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