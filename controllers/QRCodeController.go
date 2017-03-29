package controllers

import (
  "env/models"
	"github.com/astaxie/beego"
  _ "github.com/beego/i18n"
  "github.com/astaxie/beego/utils/pagination"
)

// type PageOptions struct {
//   TableName             string
//   Conditions            string
//   Currentpage           int
//   PageSize              int
//   LinkItemCount         int
//   Href                  string
//   ParamName             string
//   FirstPageText         string
//   LastPageText          string
//   PrePageText           string
//   NextPageText          string
//   EnableFirstLastLink   bool
//   EnablePreNextLink     bool
//
// }

type QRCodeController struct {
	beego.Controller
}

func (c *QRCodeController) Prepare() {
  if c.GetSession("IsLogin") == "" || c.GetSession("IsLogin") == nil {
    c.Redirect("/login", 302)
  }
  beego.Debug(c.GetSession("IsLogin"))
}

// func (this *PostsController) ListAllPosts() {
//     // sets this.Data["paginator"] with the current offset (from the url query param)
//     // 设置c.Data["paginator"]的内容（用现在的偏移？）
//     postsPerPage := 20
//     paginator := pagination.SetPaginator(this.Ctx, postsPerPage, CountPosts())
//     // fetch the next 20 posts
//     this.Data["posts"] = ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
// }

func (c *QRCodeController) Get() {
  // 默认执行的Get方法将返回所有的二维码数据
  qrlist := make([]*models.QRCode, 0)
  models.QRReadAll(&qrlist)

  codesPerPage := 1
  paginator := pagination.SetPaginator(c.Ctx, codesPerPage, models.CountCodes())

  c.Data["QRList"] = models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage)
  //c.Data["QRList"] = qrlist
  beego.Debug("找不到模板文件")
	c.TplName = "qrcode.html"
  beego.Debug("真的吗？")
}

func (c *QRCodeController) Post() {
  // 该post使用url:/plant
  // 使用flash将报错信息传到前台
  beego.Debug("post1")
  // flash := beego.NewFlash()
  // 获得表单输入
  code := models.QRCode{}
  if err := c.ParseForm(&code); err != nil {
      //handle error
  }
  if err := models.QRAddOne(&code); err != nil {
    beego.Debug(err)
  }

  beego.Debug("post2")
  c.Redirect("/code", 302)
  // id := c.Input.Get("id")
  // name := c.Input().Get("uname")
  // pic := c.Input().Get("pwd")
  // desc := c.Input().Get("desc")
  // flash.Error("密码错误")
  // flash.Store(&c.Controller)
  // c.Redirect("/admin", 302)
}
