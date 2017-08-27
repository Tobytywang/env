package back_end

import (
	"env/controllers"
	"env/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ColumnController struct {
	controllers.BaseController
}

func (c *ColumnController) Prepare() {
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "column"
}

func (c *ColumnController) Get() {
	columnlist := make([]*models.Column, 0)
	models.CReadAll(&columnlist)
	dlist := make([]*models.Column, 0)
	dlist, _ = models.SortColumn(0, columnlist, dlist)
	destlist := models.CHasSon(dlist)
	c.Data["ColumnList"] = destlist

	columnlist_column := make([]*models.Column, 0)
	models.CReadAllColumn(&columnlist_column)
	dlist_column := make([]*models.Column, 0)
	dlist_column, _ = models.SortColumn(0, columnlist_column, dlist_column)
	c.Data["ColumnList_Column"] = dlist_column
}

// 新增一个栏目
func (c *ColumnController) Add() {
	id, err := c.GetInt("id")
	if err == nil {
		beego.Debug(id)
		code, err := models.CReadById(id)
		if err == nil {
			columnlist := make([]*models.Column, 0)
			models.CReadAll(&columnlist)
			dlist := make([]*models.Column, 0)
			dlist, _ = models.SortColumn(0, columnlist, dlist)
			c.Data["ColumnList"] = dlist

			columnlist_column := make([]*models.Column, 0)
			models.CReadAllColumn(&columnlist_column)
			dlist_column := make([]*models.Column, 0)
			dlist_column, _ = models.SortColumn(0, columnlist_column, dlist_column)
			c.Data["ColumnList_Column"] = dlist_column

			c.Data["Modify"] = true
			c.Data["Column"] = code
		} else {
			beego.Debug(err)
		}
	}

	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "column_add"
}

// 新增Column
func (c *ColumnController) Post() {
	if !cpt.VerifyReq(c.Ctx.Request) {
		beego.Debug("验证未通过")
		c.Redirect("/column", 302)
	} else {
		var column models.Column
		err := c.ParseForm(&column)
		if err != nil {
			c.Redirect("/column", 302)
		}
		if column.Id != 0 {
			_, err := models.CReadById(column.Id)
			if err == nil {
				err := models.CModify(&column)
				if err != nil {
					beego.Debug(err)
				}
			}
		} else {
			err := models.CAdd(&column)
			if err != nil {
				beego.Debug(err)
			}
		}
		c.Redirect("/column", 302)
	}
}

// 修改一个栏目的内容
// 更改优先级，名称，父目录
func (c *ColumnController) Modify() {

	if !cpt.VerifyReq(c.Ctx.Request) {
		beego.Debug("验证未通过")
		c.Redirect("/column", 302)
	}
	beego.Debug("验证通过")

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
	column.Father = intfather

	err := models.CModify(column)
	beego.Debug(err)
	c.Redirect("/column", 302)
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
