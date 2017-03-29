package controllers

import (
  "env/models"
	"github.com/astaxie/beego"
  _ "github.com/beego/i18n"
  "github.com/astaxie/beego/utils/pagination"
)

type QRCodeController struct {
	beego.Controller
}

func (c *QRCodeController) Prepare() {
  if c.GetSession("IsLogin") == "" || c.GetSession("IsLogin") == nil {
    c.Redirect("/login", 302)
  }
  beego.Debug(c.GetSession("IsLogin"))
}

func (c *QRCodeController) Get() {
  // 默认执行的Get方法将返回所有的二维码数据
  qrlist := make([]*models.QRCode, 0)
  models.QRReadAll(&qrlist)

  codesPerPage := 15
  paginator := pagination.SetPaginator(c.Ctx, codesPerPage, models.CountCodes())

  beego.Debug(models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage))
  c.Data["QRList"] = models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage)
  //c.Data["QRList"] = qrlist
	c.TplName = "qrcode.html"
}

func (c *QRCodeController) Add() {
	c.TplName = "qrcode_add.html"
}

func (c *QRCodeController) Post() {
  // 该post使用url:/plant
  // 使用flash将报错信息传到前台
  beego.Debug("post1")
  // flash := beego.NewFlash()
  // 获得表单输入
  code := models.QRCode{}
  if err := c.ParseForm(&code); err != nil {
    beego.Debug(err)
  }
  if err := models.QRAddOne(&code); err != nil {
    beego.Debug(err)
  }

  beego.Debug("post2")
  c.Redirect("/code", 302)
}
