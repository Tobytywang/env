package routers

import (
	"env/controllers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/utils/captcha"
)

func init() {
	  // 登录链接
		beego.Router("/login", &controllers.LoginController{})
		// 验证码请求路径
		// beego.Handler("/captcha/*.png", captcha.Server(240, 80))
		// 首页链接
    beego.Router("/", &controllers.MainController{})
		// 二维码链接
		beego.Router("/code", &controllers.QRCodeController{})
		beego.Router("/code/add", &controllers.QRCodeController{}, "get:Add")
		beego.Router("/download", &controllers.QRCodeController{}, "get:Download")
		// 植物链接
		beego.Router("/plant", &controllers.PlantController{})
		// 文章链接
		// beego.Router("/post", &controllers.PostController{})
}
