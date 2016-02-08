package routers

import (
	"github.com/astaxie/beego"

	"github.com/nathan-osman/pratl/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/users/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/users/register", &controllers.UserController{}, "get,post:Register")
}
