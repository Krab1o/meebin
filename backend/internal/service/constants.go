package service

import "github.com/Krab1o/meebin/internal/model"

const (
	StartUtilizeCount = 0
	StartReportCount  = 0
	StartRating       = 0.0

	//TODO: make adminEmail to .env file
	AdminEmail = "admin@example.com"
)

const (
	RoleUserName  model.Role = "user"
	RoleAdminName model.Role = "admin"
)
