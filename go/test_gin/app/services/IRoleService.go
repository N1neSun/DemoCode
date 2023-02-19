package service

import "test_gin/app/models"

type IRoleService interface {
	//GetUserRoles 分页返回Articles获取用户身份信息
	GetUserRoles(userName string) []*models.Role
}
