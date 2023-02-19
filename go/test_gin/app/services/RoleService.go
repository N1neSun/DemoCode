package service

import (
	"test_gin/app/models"
	"test_gin/app/repository"
)

type RoleService struct {
	Repository repository.IRoleRepository `inject:""`
}

func NewRoleService(r repository.IRoleRepository) IRoleService {
	return &RoleService{
		Repository: r,
	}
}

func (c *RoleService) GetUserRoles(userName string) []*models.Role {
	where := models.Role{UserName: userName}
	return c.Repository.GetUserRoles(&where)
}
