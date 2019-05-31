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
	// TODO:
	return nil
}

func (c *IndexController) GetBy(id int) mvc.Result{
	// TODO:
	return nil
}

func (c *IndexController) GetSearch() mvc.Result{
	// TODO:
	return nil
}