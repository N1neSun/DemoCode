package repository

import "test_gin/app/models"

type IUserRepository interface {
	//CheckUser 身份验证
	CheckUser(where interface{}) bool
	//GetUserAvatar 获取用户头像
	GetUserAvatar(sel *string, where interface{}) *string
	//GetUserID 获取用户ID
	GetUserID(sel *string, where interface{}) int
	//GetUsers 获取用户信息
	GetUsers(PageNum int, PageSize int, total *int64, where interface{}) []*models.User
	//AddUser 新建用户
	AddUser(user *models.User) bool
	//ExistUserByName 判断用户名是否已存在
	ExistUserByName(where interface{}) bool
	//UpdateUser 更新用户
	UpdateUser(user *models.User, role *models.Role) bool
	//DeleteUser 更新用户
	DeleteUser(id int) bool
	//GetUserByID 获取用户
	GetUserByID(id int) *models.User
	//GetUserByName 获取用户的UUID
	GetUserByName(username string) string

	GetUserByWhere(where interface{}) *models.User
}
