package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
  if c.GetSession("IsLogin") == "" || c.GetSession("IsLogin") == nil {
    c.Redirect("/login", 302)
  }
}

func (c *BaseController) Get() {
	c.TplName = "bindex.html"
}
