package user

import (
	"awethome.com/action"
	"awethome.com/role"
)

type User struct {
	Attributes map[string]string
	Secret string
	Actions map[string]action.Action
	Roles map[string]role.Role
}

func (user *User) AddAction(userAction action.Action) {
	if len(user.Actions) == 0 {
		user.Actions = make(map[string]action.Action)
	}
	user.Actions[userAction.Name] =  userAction
}

func (user *User) AddRole(userRole role.Role) {
	if len(user.Roles) == 0 {
		user.Roles = make(map[string]role.Role)
	}
	user.Roles[userRole.Name] =  userRole
}