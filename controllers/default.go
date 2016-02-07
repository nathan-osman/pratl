package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"

	"github.com/nathan-osman/pratl/models"

	"errors"
)

// MainController displays the home page which provides login and registration.
type MainController struct {
	beego.Controller
}

// LoginForm provides a form for authenticating against the server.
type LoginForm struct {
	Email    string `valid:"Email;MaxSize(128)"`
	Password string `valid:"Required"`
}

// RegisterForm provides a form for account signup.
type RegisterForm struct {
	Email     string `valid:"Email"`
	Name      string `valid:"Required"`
	Password  string `valid:"Required"`
	Password2 string `valid:"Required"`
}

// Get displays a form both for logging in and registering.
func (c *MainController) Get() {
	c.Layout = "base.tpl"
	c.TplName = "index.tpl"
}

func (c *MainController) processLogin() (err error) {
	f := &LoginForm{}
	c.Data["LoginForm"] = f
	if err = c.ParseForm(f); err != nil {
		return errors.New("Unable to parse form")
	}
	v := validation.Validation{}
	b, err := v.Valid(f)
	if err != nil {
		return errors.New("Unable to validate form")
	}
	if !b {
		return errors.New("Invalid form input")
	}
	u, err := models.FindUser(f.Email)
	if err != nil {
		return errors.New("No account associated with that email address")
	}
	if err = u.Authenticate(f.Password); err != nil {
		return errors.New("Incorrect password supplied")
	}
	return
}

func (c *MainController) processRegister() (err error) {
	f := &RegisterForm{}
	c.Data["RegisterForm"] = f
	if err = c.ParseForm(f); err != nil {
		return errors.New("Unable to parse form")
	}
	v := validation.Validation{}
	b, err := v.Valid(f)
	if err != nil {
		return errors.New("Unable to validate form")
	}
	if !b {
		return errors.New("Invalid form input")
	}
	if f.Password != f.Password2 {
		return errors.New("Passwords do not match")
	}
	_, err = models.NewUser(f.Email, f.Name, f.Password)
	if err != nil {
		return errors.New("Unable to create new user")
	}
	return
}

// Post processes either the login or registration form.
func (c *MainController) Post() {
	c.Layout = "base.tpl"
	c.TplName = "index.tpl"
	switch c.GetString("Action") {
	case "login":
		if err := c.processLogin(); err != nil {
			c.Data["Error"] = err
		}
	case "register":
		if err := c.processRegister(); err != nil {
			c.Data["Error"] = err
		}
	default:
		c.Data["Error"] = errors.New("Invalid 'Action' parameter")
	}
}
