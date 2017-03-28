package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type LoginController struct {
	beego.Controller
}

var cpt *captcha.Captcha
// var store cache.cache

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 5
	cpt.StdWidth = 150
	cpt.StdHeight = 50
}
// 使用IsLogin记录登录状态

// Get逻辑
func (c *LoginController) Get() {
	flash := beego.ReadFromRequest(&c.Controller)

  // 如果输错密码
	if _, ok := flash.Data["error"]; ok {
		c.Data["Error"] = true
	}

  // 如果用户退出登录
	if c.Input().Get("exit") == "true" {
		c.SetSession("IsLogin", "")
	}

	c.TplName = "login.html"
}

// Post逻辑
func (c *LoginController) Post() {

	flash := beego.NewFlash()
	user_name := c.Input().Get("user_name")
	user_pass := c.Input().Get("user_pass")
	sys_name := beego.AppConfig.String("USER_NAME")
  sys_pass := beego.AppConfig.String("USER_PASS")
	if user_name == sys_name {
		if user_pass == sys_pass {
			if !cpt.VerifyReq(c.Ctx.Request){
				beego.Debug("验证码验证失败")
				flash.Error("验证码错误")
				flash.Store(&c.Controller)
				c.Redirect("/login", 302)
			} else {
				beego.Debug("登录成功")
				c.SetSession("IsLogin", true)
				c.Redirect("/code", 302)
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

// func (c *LoginController) VerifyCaptcha() {
// 	startTime := time.Now()
// 	captcha := c.GetString("captchaId")
// 	captchaValue := c.GetString("captcha")
// 	if !captcha.VerifyString(captchaId, captchaValue) {
// 		c.Data[""]
// 	}
// }
