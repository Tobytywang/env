package back_end

import (
  _ "os"
  _ "time"
  "strconv"
  "env/models"
  "env/controllers"
  "github.com/astaxie/beego"
  _ "github.com/beego/i18n"
  _ "github.com/astaxie/beego/utils/pagination"
)

type ColumnController struct {
  controllers.BaseController
}

func (c *ColumnController) Prepare() {
	beego.Debug("column")
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "column"
}

// 显示所有栏目
func (c *ColumnController) Get() {
	columnlist := make([]*models.Column, 0)
	models.CReadAll(&columnlist)
	dlist := make([]*models.Column, 0)
	dlist, _ = models.SortColumn(0, columnlist, dlist)

	c.Data["ColumnList"] = dlist
	beego.Debug("中")
	c.Data["Type"] = "index" 
}

// 新增一个栏目
func (c *ColumnController) Add() {
	if(c.Ctx.Input.IsPost() == true) {
		column := models.Column{}
		c.ParseForm(&column)
		ctype := c.Input().Get("type")
		column.Type = ctype
		models.CAdd(&column)
		c.Redirect("/column", 302)
	} else {
		c.TplName = "back_end/public.html"
		c.Data["Tpl"] = "column_add"
	}
}

// 删除一个栏目
func (c *ColumnController) Del() {

	id := c.Input().Get("id")
	intid, _ := strconv.Atoi(id)
	columnlist := make([]*models.Column, 0)
	models.CReadAll(&columnlist)
	dlist := make([]*models.Column, 0)
	dlist, columnlist = models.CFindSon(intid, columnlist, dlist)
	if len(dlist) > 0 {
		c.Redirect("/column", 302)
	} else {
		models.CDel(intid)
		c.Redirect("/column", 302)
	}
}

// 查找一篇文章
func (c *ColumnController) Search() {
  // content := c.GetString("content")
  // beego.Debug(content)
  // qrlist := models.QRSearch(content)
  // beego.Debug(qrlist)
  // c.Data["QRList"] = qrlist
  // // c.TplName = "back_end/qrcode.html"
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "post_add"
}

// 修改一个栏目的内容
func (c *ColumnController) Modify() {

	pri := c.Input().Get("pri")
	intpri, _ := strconv.Atoi(pri)

	id := c.Input().Get("id")
	intid, _ := strconv.Atoi(id)

	column := new(models.Column)
	column = models.CReadById(intid)
	column.Pri = intpri
	models.CModify(column)
	c.Redirect("/column", 302)
}
