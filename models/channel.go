package models

import (
	"github.com/astaxie/beego/orm"
)

// Channel represents an individual "room" with messages. Room owners have
// "administration" abilities and can moderate messages.
type Channel struct {
	Id          int64   `orm:"auto"`
	Name        string  `orm:"size(32)"`
	Description string  `orm:"size(128)"`
	Owners      []*User `orm:"rel(m2m)"`
}

func init() {
	orm.RegisterModel(new(Channel))
}

// NewChannel creates a new channel with the specified owner.
func NewChannel(name, description string, owner *User) (c *Channel, err error) {
	c = &Channel{
		Name:        name,
		Description: description,
		Owners: []*User{
			owner,
		},
	}
	_, err = orm.NewOrm().Insert(c)
	return
}
