package role

import (
	"awethome/server/action"
	"awethome/server/resource"
)

type Role struct {
	Name string
	Actions map[string]action.Action
	Resources map[string]resource.Resource
}

func (role *Role) AddAction(roleAction action.Action) {
	if len(role.Actions) == 0 {
		role.Actions = make(map[string]action.Action)
	}
	role.Actions[roleAction.Name] =  roleAction
}

func (role *Role) AddResource(roleResource resource.Resource) {
	if len(role.Resources) == 0 {
		role.Resources = make(map[string]resource.Resource)
	}
	role.Resources[roleResource.Name] =  roleResource
}