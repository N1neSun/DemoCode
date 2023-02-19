package repository

import "test_gin/app/models"

type IRoleRepository interface {
	GetUserRoles(where interface{}) []*models.Role

	GetRoles(sel *string, where interface{}) []string

	AddRole(role *models.Role) bool

	GetRole(where interface{}) *models.Role
}
