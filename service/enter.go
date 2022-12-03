package service

type group struct {
	AdminUserService AdminUserService
	UserService      UserService
}

var Group = new(group)
