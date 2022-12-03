package router

type group struct {
	AdminRouterGroup AdminRouterGroup
	UserRouterGroup  UserRouterGroup
}

var Group = new(group)
