package back_end

import (
	"env/controllers"
	"env/models"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type QRCodeController struct {
	controllers.BaseController
}

// Get方法查看所有的二维码
func (c *QRCodeController) Get() {
	codesPerPage := 15
	paginator := pagination.SetPaginator(c.Ctx, codesPerPage, models.CountCodes())

	sort := c.GetString("sort")
	beego.Debug(sort)
	if (sort != "readup") && (sort != "readdown") && (sort != "iddown") {
		sort = "id"
	}
	c.Data["URL"] = beego.AppConfig.String("WEB_URL")
	c.Data["QRList"] = models.ListCodesByOffsetAndLimit(sort, paginator.Offset(), codesPerPage)
	// beego.Debug(models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage))

	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "qrcode"
}

// Post增加
func (c *QRCodeController) Post() {
	var code models.QRCode
	if err := c.ParseForm(&code); err != nil {
		beego.Debug(err)
		c.Redirect("/code", 302)
	}
	filetype := c.GetString("filetype")
	beego.Debug(code)
	if code.Id != 0 {
		beego.Debug("id不为空")
		beego.Debug(code.Id)
		if data, err := models.QRReadById(code.Id); err == nil {
			beego.Debug(data)
			beego.Debug(code)
			if _, err := c.SaveFile(&code, filetype); err != nil {
				beego.Debug(err)
			}
			beego.Debug(code)
			if err := models.QRUpdate(&code); err != nil {
				beego.Debug(err)
			}
			beego.Debug("更新成功")
		}
	} else {
		beego.Debug("id为空")
		beego.Debug(code.Id)
		beego.Debug("存储图片<")
		if _, err := c.SaveFile(&code, filetype); err != nil {
			beego.Debug(err)
		}
		beego.Debug("存储图片>")
		if err := models.QRAddOne(&code); err != nil {
			beego.Debug(err)
		}
	}
	c.Redirect("/code", 302)
}

// Add方法增加一个二维码
func (c *QRCodeController) Add() {
	if id, err := c.GetInt("id"); err == nil {
		if code, err := models.QRReadById(id); err == nil {
			beego.Debug(code)
			c.Data["Modify"] = true
			c.Data["Code"] = code
		}
	}

	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "qrcode_add"
}

// Download方法下载一个植物的二维码
func (c *QRCodeController) Download() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Debug(err)
	}

	code, err := models.QRReadById(id)
	beego.Debug(code)
	if err != nil {
		beego.Debug(err)
	}

	code_string, err := models.Create_qrcode(code, code.Name)
	c.Ctx.Output.Download(code_string)
	os.Remove(code_string)
	c.Redirect("/code", 302)
}

// Del删除一个二维码
func (c *QRCodeController) Del() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Debug(err)
	}
	err = models.QRDel(id)
	if err != nil {
		beego.Debug(err)
	}
	c.Redirect("/code", 302)
}

// Search根据内容筛选二维码
func (c *QRCodeController) Search() {
	content := c.GetString("content")

	qrlist := models.QRSearch(content)

	c.Data["QRList"] = qrlist

	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "qrcode"
}

// 存储上传的图片
func (c *QRCodeController) SaveFile(p *models.QRCode, filetype string) (string, error) {
	filepath := "static/upload/"
	p.Pic = filepath + "/" + p.Name + filetype
	_, _, err := c.GetFile("pic")
	if err == nil {
		os.MkdirAll(filepath, 0777)
		if err := c.SaveToFile("pic", filepath+"/"+p.Name+filetype); err != nil {
			beego.Debug(err)
		}
	} else {
		return "", err
	}
	return p.Pic, nil
}
