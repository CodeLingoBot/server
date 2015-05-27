package user

import (
	"awethome.com/action"
	"awethome.com/role"
)

type User struct {
	Attributes map[string]string
	Secret string
	Actions []action.Action
	Roles []role.Role
}

func (user *User) AddAction(action action.Action) {
	user.Actions = append(user.Actions, action)
}

func (user *User) AddRole(role role.Role) {
	user.Roles = append(user.Roles, role)
}