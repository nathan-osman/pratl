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
