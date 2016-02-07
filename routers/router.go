package routers

import (
	"github.com/astaxie/beego"

	"github.com/nathan-osman/pratl/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
