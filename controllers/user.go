package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"

	"github.com/nathan-osman/pratl/models"
)

// UserController facilitates user registration and authentication.
type UserController struct {
	beego.Controller
}

// LoginForm provides a form for user authentication.
type LoginForm struct {
	Email    string `valid:"Email;MaxSize(128)"`
	Password string `valid:"Required"`
}

// Login attempts to authenticate a user. If successful, they are redirected to
// the list of channels that they can join.
func (c *UserController) Login() {
	f := &LoginForm{}
	if c.Ctx.Request.Method == "POST" {
		if err := c.ParseForm(f); err == nil {
			v := validation.Validation{}
			b, err := v.Valid(f)
			if b && err == nil {
				u, err := models.FindUser(f.Email)
				if err == nil {
					if err := u.Authenticate(f.Password); err == nil {
						c.SetSession("user_id", u.Id)
						r := c.GetString("redirect")
						if r == "" {
							r = c.URLFor("ChatController.Index")
						}
						c.Redirect(r, 302)
						return
					} else {
						c.Data["Error"] = "Invalid password"
					}
				} else {
					c.Data["Error"] = "No account with that email address"
				}
			} else {
				c.Data["Error"] = "Invalid form input"
			}
		} else {
			c.Data["Error"] = "Unable to parse form"
		}
	}
	c.TplName = "user/login.tpl"
	c.Render()
}

// Logout destroys a user's session and returns them to the home page.
func (c *UserController) Logout() {
	c.DelSession("user_id")
	c.Redirect(c.URLFor("IndexController.Get"), 302)
}

// RegisterForm provides a form for account signup.
type RegisterForm struct {
	Email     string `valid:"Email;MaxSize(128)"`
	Name      string `valid:"Required;MaxSize(128)"`
	Password  string `valid:"Required"`
	Password2 string `valid:"Required"`
}

// Register displays the registration form.
func (c *UserController) Register() {
	f := &RegisterForm{}
	if c.Ctx.Request.Method == "POST" {
		if err := c.ParseForm(f); err == nil {
			v := validation.Validation{}
			b, err := v.Valid(f)
			if b && err == nil {
				if f.Password == f.Password2 {
					u, err := models.NewUser(f.Email, f.Name, f.Password)
					if err == nil {
						c.SetSession("user_id", u.Id)
						c.Redirect(c.URLFor("ChatController.Index"), 302)
						return
					} else {
						c.Data["Error"] = "Unable to complete registration"
					}
				} else {
					c.Data["Error"] = "Passwords don't match"
				}
			} else {
				c.Data["Error"] = "Invalid form input"
			}
		} else {
			c.Data["Error"] = "Unable to parse form"
		}
	}
	c.TplName = "user/register.tpl"
	c.Render()
}
