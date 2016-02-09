package controllers

import (
	"github.com/astaxie/beego"
)

// IndexController displays the home page.
type IndexController struct {
	beego.Controller
}

// Get displays the home page with links for logging in and registering. If the
// user is already logged in, they are redirected to the channels page.
func (c *IndexController) Get() {
	if c.GetSession("user_id") == nil {
		c.TplName = "index.tpl"
		c.Render()
	} else {
		c.Redirect("/channels", 302)
	}
}
