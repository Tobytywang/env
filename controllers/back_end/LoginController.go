// 用户登录验证
package back_end

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type LoginController struct {
	beego.Controller
}

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 5
	cpt.StdWidth = 150
	cpt.StdHeight = 50
}

// Get方法（获取登录页面）
func (c *LoginController) Get() {

	// 不需要密码
	// c.SetSession("IsLogin", true)
	// c.Redirect("/admin", 302)

	flash := beego.ReadFromRequest(&c.Controller)

	// 如果输错密码
	if _, ok := flash.Data["error"]; ok {
		c.Data["Error"] = true
	}

	// 如果用户退出登录
	if c.Input().Get("exit") == "true" {
		c.SetSession("IsLogin", "")
	}

	c.TplName = "back_end/login.html"
}

// Post方法（提交登录表单）
func (c *LoginController) Post() {

	flash := beego.NewFlash()
	user_name := c.Input().Get("user_name")
	user_pass := c.Input().Get("user_pass")
	sys_name := beego.AppConfig.String("USER_NAME")
	sys_pass := beego.AppConfig.String("USER_PASS")
	if user_name == sys_name {
		if user_pass == sys_pass {
			if !cpt.VerifyReq(c.Ctx.Request) {
				beego.Debug("验证码验证失败")
				flash.Error("验证码错误")
				flash.Store(&c.Controller)
				c.Redirect("/login", 302)
			} else {
				beego.Debug("登录成功")
				c.SetSession("IsLogin", true)
				c.Redirect("/admin", 302)
			}
		} else {
			beego.Debug("密码错误")
			flash.Error("密码错误")
			flash.Store(&c.Controller)
			c.Redirect("/login", 302)
		}
	} else {
		beego.Debug("账户名错误")
		flash.Error("账户名错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}
