package service

import (
	"test_gin/app/domain"
	"test_gin/app/models"
	"test_gin/app/repository"
	"test_gin/common/logger"
)

type UserService struct {
	Repository     repository.IUserRepository `inject:""`
	RoleRepository repository.IRoleRepository `inject:""`
	Log            logger.ILogger             `inject:""`
}

func NewUserService(userRepository repository.IUserRepository, roleRepository repository.IRoleRepository, log *logger.Logger) IUserService {
	return &UserService{
		Repository:     userRepository,
		RoleRepository: roleRepository,
		Log:            log,
	}
}

func (a *UserService) CheckUser(username string, password string) bool {
	where := models.User{Username: username, Password: password}
	return a.Repository.CheckUser(where)
}

func (a *UserService) GetUserAvatar(username string) *string {
	where := models.User{Username: username}
	sel := "avatar"
	return a.Repository.GetUserAvatar(&sel, &where)
}

func (a *UserService) GetRoles(username string) []string {
	userWhere := models.User{Username: username}
	userSel := "id"
	userID := a.Repository.GetUserID(&userSel, &userWhere)
	roleWhere := models.Role{UserID: userID}
	roleSel := "value"
	return a.RoleRepository.GetRoles(&roleSel, &roleWhere)
}

func (a *UserService) GetUsers(page, pagesize int, maps interface{}) interface{} {
	res := make(map[string]interface{}, 2)
	var total int64
	users := a.Repository.GetUsers(page, pagesize, &total, maps)
	var pageUsers []domain.Users
	for _, user := range users {
		var pageUser domain.Users
		pageUser.ID = user.ID
		pageUser.Name = user.Username
		pageUser.Password = user.Password
		pageUser.Avatar = user.Avatar
		pageUser.CreatedOn = user.CreatedOn.Format("2006-01-02 15:04:05")
		pageUsers = append(pageUsers, pageUser)
	}
	res["list"] = pageUsers
	res["total"] = total
	return &res
}

func (a *UserService) AddUser(user *models.User) bool {
	isOK := a.Repository.AddUser(user)
	if !isOK {
		return false
	}
	var role models.Role
	role.UserID = user.ID
	role.UserName = user.Username
	role.Value = "normal"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	isOK = a.RoleRepository.AddRole(&role)
	if isOK {
		return true
	}
	return a.Repository.DeleteUser(user.ID)
}

func (a *UserService) ExistUserByName(username string) bool {
	where := models.User{Username: username}
	return a.Repository.ExistUserByName(&where)
}

func (a *UserService) UpdateUser(modUser *models.User) bool {
	user := a.Repository.GetUserByID(modUser.ID)
	//不允许更新用户名
	// user.Username = modUser.Username
	user.Password = modUser.Password
	user.ModifiedBy = modUser.ModifiedBy
	user.UserType = modUser.UserType
	roleWhere := models.Role{UserID: user.ID}
	role := a.RoleRepository.GetRole(&roleWhere)
	// role.UserName = user.Username
	role.Value = "test"
	if user.UserType == 1 {
		role.Value = "admin"
	}
	return a.Repository.UpdateUser(user, role)
}

func (a *UserService) DeleteUser(id int) bool {
	user := a.Repository.GetUserByID(id)
	if user.Username == "admin" {
		a.Log.Errorf("删除用户失败:不能删除admin账号")
		return false
	}
	return a.DeleteUser(id)
}
