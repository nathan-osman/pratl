package models

import (
	"github.com/astaxie/beego/orm"

	"time"
)

// Message represents an individual chat message.
type Message struct {
	Id      int64  `orm:"auto"`
	Body    string `orm:"size(200)"`
	Time    time.Time
	Author  *User    `orm:"rel(fk)"`
	Channel *Channel `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Message))
}
