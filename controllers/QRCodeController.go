package controllers

import (
  "os"
  "time"
  "strconv"
  "math/rand"
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
  // beego.Debug(c.GetSession("IsLogin"))
}

func (c *QRCodeController) Get() {
  // 默认执行的Get方法将返回所有的二维码数据
  qrlist := make([]*models.QRCode, 0)
  models.QRReadAll(&qrlist)

  codesPerPage := 15
  paginator := pagination.SetPaginator(c.Ctx, codesPerPage, models.CountCodes())

  // beego.Debug(models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage))
  c.Data["QRList"] = models.ListCodesByOffsetAndLimit(paginator.Offset(), codesPerPage)
  //c.Data["QRList"] = qrlist
	c.TplName = "qrcode.html"
}

func (c *QRCodeController) Add() {
	c.TplName = "qrcode_add.html"
}

func (c *QRCodeController) Download() {
  id, err := c.GetInt("id");
  if err != nil {
    beego.Debug(err)
  }
  code, err := models.QRReadById(id);
  beego.Debug(code)
  if err != nil {
    beego.Debug(err)
  }
  c.Ctx.Output.Download(code.Code);
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
  ttype := c.Input().Get("filetype")
  beego.Debug(ttype)
  if _, err := c.SaveFile(&code, ttype); err != nil {
    beego.Debug(err)
  }
  if err := models.QRAddOne(&code); err != nil {
    beego.Debug(err)
  }

  // beego.Debug("post2")
  c.Redirect("/code", 302)
}

// 处理上传文件
// 返回文件的路径
func (c *QRCodeController) SaveFile(p *models.QRCode, ttype string) (string, error) {
	// 重命名为随机数
	p.Name = c.RandName() + ttype
	filepath := "static/upload/" + time.Now().Format("2006-01-02")
	p.Pic = filepath + "/" + p.Name
	// 存入文件系统
	beego.Debug(p)
	_, _, err := c.GetFile("pic")
  beego.Debug(err)
  beego.Debug("获取到了文件")
	if err == nil {
		os.MkdirAll(filepath, 0777)
		if err:=c.SaveToFile("pic", filepath+"/"+p.Name); err!=nil{
      beego.Debug(err)
    }
	} else {
		return "", err
	}
	// 存入数据库
	// models.DAdd(p)
	// beego.Debug(p)
	return p.Pic, nil
}

func (c *QRCodeController) RandName() string {
	var name string
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		temp := r.Intn(9)
		name = name + strconv.Itoa(temp)
	}
	return name
}
