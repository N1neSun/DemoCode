package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `json:"created_on"`
	ModifiedOn time.Time `json:"modified_on"`
	Username   string    `json:"username" validate:"required"`
	Useruuid   string    `json:"useruuid"`
	Password   string    `json:"password" validate:"required"`
	LastLogin  time.Time `json:"lastlogin"`
	Avatar     string    `json:"avatar"`
	UserType   int       `json:"user_type"`
	Deleted    int       `json:"deteled"`
	State      int       `json:"state"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
}

//BeforeCreate CreatedOn赋值
func (user *User) BeforeCreate(db *gorm.DB) error {
	user.CreatedOn = time.Now()
	u4, _ := uuid.NewV4()
	user.Useruuid = u4.String()
	return nil
}

//BeforeUpdate ModifiedOn赋值
func (user *User) BeforeUpdate(db *gorm.DB) error {
	user.ModifiedOn = time.Now()
	return nil
}

// UserRole 用户身份结构体
type UserRole struct {
	UserName  string
	UserRoles []*Role
}
