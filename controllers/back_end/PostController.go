package back_end

import (
	"env/controllers"
	"github.com/astaxie/beego"
)

type PostController struct {
  controllers.BaseController
}


// Get方法查看所有的文章
func (c *PostController) Get() {

  // qrlist := make([]*models.QRCode, 0)
  // models.QRReadAll(&qrlist)

  // codesPerPage := 15
  // paginator := pagination.SetPaginator(c.Ctx, codesPerPage, models.CountCodes())

  // c.Data["URL"] = beego.AppConfig.String("WEB_URL")
  // c.Data["QRList"] = models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage)
  // beego.Debug(models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage))
  // // c.TplName = "back_end/qrcode.html"
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "post"
}

func (c *PostController) Post() {
  var content = c.GetString("content")
  beego.Debug(content)
  c.TplName = "index.tpl"
}

// 增加一篇文章
func (c *PostController) Add() {
  // if id, err := c.GetInt("id"); err == nil{
  //   if code, err := models.QRReadById(id); err == nil{
  //     beego.Debug(code)
  //     c.Data["Modify"] = true
  //     c.Data["Code"] = code
  //   }
  // }
  // // c.TplName = "back_end/qrcode_add.html"
  ////////////////////////////////////////////////////
  // if(c.Ctx.Input.Is("POST") == true) {

    // c.TplName = "index.tpl"
    c.TplName = "back_end/public.html"
    c.Data["Tpl"] = "post_add"
}

// 删除一篇文章
func (c *PostController) Del() {
  // id, err := c.GetInt("id");
  // if err != nil {
  //   beego.Debug(err)
  // }
  // err = models.QRDel(id);
  // if err != nil {
  //   beego.Debug(err)
  // }
  c.Redirect("/post", 302)
}


// 查找一篇文章
func (c *PostController) Search() {
  // content := c.GetString("content")
  // beego.Debug(content)
  // qrlist := models.QRSearch(content)
  // beego.Debug(qrlist)
  // c.Data["QRList"] = qrlist
  // // c.TplName = "back_end/qrcode.html"
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "post_add"
}