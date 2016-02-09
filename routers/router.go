package routers

import (
	"github.com/astaxie/beego"

	"github.com/nathan-osman/pratl/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/channels", &controllers.ChatController{}, "get:Index")
	beego.Router("/channels/new", &controllers.ChatController{}, "get,post:New")
	beego.Router("/channels/:id:int/.*", &controllers.ChatController{}, "get:View")
	beego.Router("/users/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/users/register", &controllers.UserController{}, "get,post:Register")
}
