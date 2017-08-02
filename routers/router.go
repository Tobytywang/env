// 记录程序的路由信息
package routers

import (
	"env/controllers"
	"github.com/astaxie/beego"
	"env/controllers/back_end"
	"env/controllers/front_end"
)

func init() {
	/////////////////////////////////////////////////////////////
	// 前台链接
	/////////////////////////////////////////////////////////////
	beego.Router("/", &controllers.MainController{})
	beego.Router("/plant", &front_end.PlantController{})

	/////////////////////////////////////////////////////////////
	// 后台链接
	/////////////////////////////////////////////////////////////
	// 登录链接
	beego.Router("/login", &back_end.LoginController{})
	// 首页链接
	beego.Router("/admin", &controllers.BaseController{})
	// 栏目链接
	beego.Router("/column", &back_end.ColumnController{})
	beego.Router("/column/add", &back_end.ColumnController{}, "get:Add")
	beego.Router("/column/del", &back_end.ColumnController{}, "get:Del")
	beego.Router("/column/search", &back_end.ColumnController{}, "get,post:Search")
	// 二维码链接
	beego.Router("/code", &back_end.QRCodeController{})
	beego.Router("/code/add", &back_end.QRCodeController{}, "get:Add")
	beego.Router("/code/del", &back_end.QRCodeController{}, "get:Del")
	beego.Router("/code/search", &back_end.QRCodeController{}, "get,post:Search")
	beego.Router("/download", &back_end.QRCodeController{}, "get:Download")
	beego.AutoRouter(&back_end.QRCodeController{})
	// 文章链接
	beego.Router("/post", &back_end.PostController{})
	beego.Router("/post/add", &back_end.PostController{}, "get,post:Add")
	beego.Router("/post/del", &back_end.PostController{}, "get:Del")
	beego.Router("/post/search", &back_end.PostController{}, "get,post:Search")
	// 设置链接
	beego.Router("/setting", &back_end.SettingController{})
}
