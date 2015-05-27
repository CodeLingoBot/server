package user

import (
	"awethome/server/action"
	"awethome/server/role"
	"awethome/server/resource"
)

type User struct {
	Id         string
	Secret     string
	Attributes map[string]string
	Actions    map[string]action.Action
	Roles      map[string]role.Role
	Resources  map[string]resource.Resource
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

func (user *User) AddResource(userResource resource.Resource) {
	if len(user.Resources) == 0 {
		user.Resources = make(map[string]resource.Resource)
	}
	user.Resources[userResource.Name] =  userResource
}