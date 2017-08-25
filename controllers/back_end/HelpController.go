// 用户登录验证
package back_end

import (
	"env/controllers"
)

type HelpController struct {
	controllers.BaseController
}

func (c *HelpController) Get() {
	c.TplName = "back_end/public.html"
	c.Data["Tpl"] = "help"
}
