package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"

	"github.com/nathan-osman/pratl/controllers"
	"github.com/nathan-osman/pratl/models"

	"fmt"
	"net/url"
)

// FilterUser determines if the user is logged in.
func FilterUser(ctx *context.Context) {
	if id, ok := ctx.Input.Session("user_id").(int64); ok {
		if u, err := models.GetUser(id); err == nil {
			ctx.Input.SetData("user", u)
		}
	}
}

// FilterAuth prevents the user from accessing protected pages if they are not
// logged in.
func FilterAuth(ctx *context.Context) {
	if ctx.Input.GetData("user") == nil {
		ctx.Redirect(302, fmt.Sprintf(
			"%s?redirect=%s",
			beego.URLFor("UserController.Login"),
			url.QueryEscape(ctx.Request.URL.Path),
		))
	}
}

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/channels", &controllers.ChatController{}, "get:Index")
	beego.Router("/channels/new", &controllers.ChatController{}, "get,post:New")
	beego.Router("/channels/:id:int", &controllers.ChatController{}, "get:View")
	beego.Router("/users/login", &controllers.UserController{}, "get,post:Login")
	beego.Router("/users/logout", &controllers.UserController{}, "get,post:Logout")
	beego.Router("/users/register", &controllers.UserController{}, "get,post:Register")

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/channels/*", beego.BeforeRouter, FilterAuth)
}
