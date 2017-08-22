package back_end

import (
	"env/controllers"
	"github.com/astaxie/beego"
)

type SettingController struct {
  controllers.BaseController
}


// Get方法查看设置项
func (c *SettingController) Get() {

  c.Data["username"] = beego.AppConfig.String("USER_NAME")
  c.Data["password"] = beego.AppConfig.String("USER_PASS")
  // c.Data["db_type"] = beego.AppConfig.String("DB_TYPE")
  c.Data["db_name"] = beego.AppConfig.String("DB_NAME")
  c.Data["db_user"] = beego.AppConfig.String("DB_USER")
  c.Data["db_pass"] = beego.AppConfig.String("DB_PASSWD")
  c.Data["web_url"] = beego.AppConfig.String("WEB_URL")
  // c.Data["qr_path"] = beego.AppConfig.String("QRPATH")

  c.Data["Modify"] = true
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "setting"
}

// Post方法修改设置项目
func (c *SettingController) Post() {

  // username := c.GetString("username")
  // password := c.GetString("password")
  // db_name := c.GetString("db_name")
  // db_user := c.GetString("db_user")
  // db_pass := c.GetString("db_pass")
  // web_url := c.GetString("web_url")

  
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "setting"
}
