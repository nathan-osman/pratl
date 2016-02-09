package controllers

import (
	"github.com/astaxie/beego"
)

// ChatController displays the channel list and individual channel pages.
type ChatController struct {
	beego.Controller
}

// Index displays the list of channels.
func (c *ChatController) Index() {
	c.TplName = "chat/index.tpl"
	c.Render()
}

// NewChannelForm provides a form for creating a new channel.
type NewChannelForm struct {
	Name        string `valid:"Required;MaxSize(32)"`
	Description string `valid:"Required;MaxSize(128)"`
}

// New enables users to create new channels.
func (c *ChatController) New() {
	c.TplName = "chat/new.tpl"
	c.Render()
}

// View displays the page for a channel.
func (c *ChatController) View() {
	c.TplName = "chat/view.tpl"
	c.Render()
}
