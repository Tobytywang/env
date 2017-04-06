// 手机端查看植物信息的页面
package controllers

import (
  "env/models"
	"github.com/astaxie/beego"
)

type PlantController struct {
	beego.Controller
}

// 获取根据Id填充页面内容
func (c *PlantController) Get() {
	id, err := c.GetInt("id");
  if err != nil {
    beego.Debug(err)

  }
	code, err := models.QRReadById(id)
	// beego.Debug(code)
	if err != nil {
		beego.Debug(err)
    c.Abort("404")
	}
  beego.Debug("增加阅读数")
  models.QRRead(id)
	c.Data["Plant"] = code
	c.TplName = "plant.html"
}
