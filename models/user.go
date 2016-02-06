package models

import (
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"

	"encoding/base64"
)

// User represents a registered user. Passwords are stored with bcrypt which
// uses a salt and computationally-expensive hashing algorithm.
type User struct {
	Id       int64  `orm:"auto"`
	Username string `orm:"size(32);index;unique"`
	Password string `orm:"size(128)"`
	Email    string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// NewUser creates a new user and updates the database.
func NewUser(username, password, email string) (u *User, err error) {
	u = &User{
		Username: username,
		Email:    email,
	}
	if err = u.SetPassword(password); err != nil {
		return
	}
	_, err = orm.NewOrm().Insert(u)
	return
}

// FindUser attempts to retrieve a user by their name.
func FindUser(username string) (u *User, err error) {
	o := orm.NewOrm()
	u = &User{Username: username}
	err = o.Read(u)
	return
}

// SetPassword assigns a new password to the user. Note that this does NOT
// update the database, so that must be done separately.
func (u *User) SetPassword(password string) (err error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return
	}
	u.Password = base64.StdEncoding.EncodeToString(h)
	return
}

// Authenticate attempts to authenticate the user.
func (u *User) Authenticate(password string) (err error) {
	s, err := base64.StdEncoding.DecodeString(u.Password)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword(s, []byte(password))
	return
}

// TODO: options for updating and deleting users
