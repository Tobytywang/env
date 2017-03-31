package controllers

import (
  "env/models"
	"github.com/astaxie/beego"
)

type PlantController struct {
	beego.Controller
}

func (c *PlantController) Prepare() {
	if c.GetSession("IsLogin") == "" || c.GetSession("IsLogin") == nil {
		c.Redirect("/login", 302)
	}
}
func (c *PlantController) Get() {
  // 默认执行的Get方法将返回所有的二维码数据
  // qrlist := make([]*models.QRCode, 0)
  // models.QRReadAll(&qrlist)

  // c.Data["QRList"] = qrlist
	id, err := c.GetInt("id");
  if err != nil {
    beego.Debug(err)
  }
	code, err := models.QRReadById(id);
	beego.Debug(code)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["Plant"] = code
	c.TplName = "plant.html"
}
//
// func (c *PlantController) Post() {
//   // 该post使用url:/plant
//   // 使用flash将报错信息传到前台
//   beego.Debug("post1")
//   // flash := beego.NewFlash()
//   // 获得表单输入
//   code := models.QRCode{}
//   if err := c.ParseForm(&code); err != nil {
//       //handle error
//   }
//   if err := models.QRAddOne(&code); err != nil {
//     beego.Debug(err)
//   }
//
//   beego.Debug("post2")
//   c.Redirect("/plant", 302)
//   // id := c.Input.Get("id")
//   // name := c.Input().Get("uname")
//   // pic := c.Input().Get("pwd")
//   // desc := c.Input().Get("desc")
//   // flash.Error("密码错误")
//   // flash.Store(&c.Controller)
//   // c.Redirect("/admin", 302)
// }
