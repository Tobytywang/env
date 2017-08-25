package back_end

import (
	"env/controllers"
	"env/models"
	_ "os"
	"strconv"
	_ "time"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/utils/pagination"
	_ "github.com/beego/i18n"
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
}

// 新增Column
func (c *ColumnController) Post() {
	var column models.Column
	if err := c.ParseForm(&column); err != nil {
		beego.Debug(err)
		// 说明有错误，跳转到查看所有项目界面？
		// 使用flash提示
		c.Redirect("/column", 302)
	}
	if column.Id != 0 {
		// 有Id，表示修改
		if _, err := models.CReadById(column.Id); err == nil {
			if err := models.CModify(&column); err != nil {
				beego.Debug(err)
			}
		}
	} else {
		// 没有Id，表示新增
		if err := models.CAdd(&column); err != nil {
			beego.Debug(err)
		}
	}
	c.Redirect("/column", 302)
}

// 新增一个栏目
func (c *ColumnController) Add() {
	if id, err := c.GetInt("id"); err == nil {
		beego.Debug(id)
		if code, err := models.CReadById(id); err == nil {
			beego.Debug(code)
			// c.Data["Modify"] = true

			columnlist := make([]*models.Column, 0)
			models.CReadAll(&columnlist)
			dlist := make([]*models.Column, 0)
			dlist, _ = models.SortColumn(0, columnlist, dlist)
			c.Data["ColumnList"] = dlist

			c.Data["Modify"] = true
			c.Data["Column"] = code
		} else {
			beego.Debug(err)
		}
	}
	// c.TplName = "back_end/qrcode_add.html"
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "column_add"
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
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "post_add"
}

// 修改一个栏目的内容
// 更改优先级，名称，父目录
func (c *ColumnController) Modify() {

	id := c.Input().Get("id")
	intid, _ := strconv.Atoi(id)

	// pri := c.Input().Get("pri")
	// intpri, _ := strconv.Atoi(pri)

	name := c.Input().Get("name")

	father := c.Input().Get("father")
	intfather, _ := strconv.Atoi(father)

	column := new(models.Column)
	column, _ = models.CReadById(intid)
	column.Id = intid
	column.Name = name
	// column.Pri = intpri
	column.Father = intfather

	beego.Debug(column)
	beego.Debug(intid)
	beego.Debug(name)
	// beego.Debug(intpri)
	beego.Debug(intfather)
	err := models.CModify(column)
	beego.Debug(err)
	c.Redirect("/column", 302)
}
