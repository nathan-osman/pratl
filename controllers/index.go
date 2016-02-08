package controllers

import (
	"github.com/astaxie/beego"
)

// IndexController displays the home page.
type IndexController struct {
	beego.Controller
}

// Get displays the home page, which provides a form for logging in.
func (c *IndexController) Get() {
	c.TplName = "index.tpl"
}
