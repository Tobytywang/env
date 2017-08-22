// 手机端查看植物信息的页面
package front_end

import (
	_ "env/models"
	"github.com/astaxie/beego"
)

type ColumnController struct {
	beego.Controller
}

// 获取根据Id填充页面内容
func (c *ColumnController) Get() {
	// id, err := c.GetInt("id");
	// if err != nil {
	// 	beego.Debug(err)
	// }
	// code, err := models.CReadById(id)
	// // beego.Debug(code)
	// if err != nil {
	// 	beego.Debug(err)
	// c.Abort("404")
	// }
	// c.Data["Plant"] = code
	c.TplName = "front_end/plant.html"
}
