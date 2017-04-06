// 记录程序的路由信息
package routers

import (
	"env/controllers"
	"github.com/astaxie/beego"
)

func init() {
	  // 登录链接
		beego.Router("/login", &controllers.LoginController{})
		// 首页链接
    beego.Router("/", &controllers.MainController{})
		beego.Router("/admin", &controllers.BaseController{})
		// 二维码链接
		beego.Router("/code", &controllers.QRCodeController{})
		beego.Router("/download", &controllers.QRCodeController{}, "get:Download")
		beego.AutoRouter(&controllers.QRCodeController{})
		beego.Router("/code/add", &controllers.QRCodeController{}, "get:Add")
		beego.Router("/code/del", &controllers.QRCodeController{}, "get:Del")
		beego.Router("/code/search", &controllers.QRCodeController{}, "get,post:Search")
		// 植物链接
		beego.Router("/plant", &controllers.PlantController{})
		// 文章链接
		// beego.Router("/post", &controllers.PostController{})
}
