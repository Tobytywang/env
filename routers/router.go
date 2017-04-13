// 记录程序的路由信息
package routers

import (
	"env/controllers"
	"github.com/astaxie/beego"
	"env/controllers/back_end"
	"env/controllers/front_end"
)

func init() {
	// 登录链接
	beego.Router("/login", &back_end.LoginController{})
	// 首页链接
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.BaseController{})
	// 二维码链接
	beego.Router("/code", &back_end.QRCodeController{})
	beego.Router("/download", &back_end.QRCodeController{}, "get:Download")
	beego.AutoRouter(&back_end.QRCodeController{})
	beego.Router("/code/add", &back_end.QRCodeController{}, "get:Add")
	beego.Router("/code/del", &back_end.QRCodeController{}, "get:Del")
	beego.Router("/code/search", &back_end.QRCodeController{}, "get,post:Search")
	// 植物链接
	beego.Router("/plant", &front_end.PlantController{})
	// 文章链接
	// beego.Router("/post", &controllers.PostController{})
}
