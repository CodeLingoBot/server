package role

import (
	"awethome.com/action"
)

type Role struct {
	Name string
	Actions map[string]action.Action
}

func (role *Role) AddAction(roleAction action.Action) {
	if len(role.Actions) == 0 {
		role.Actions = make(map[string]action.Action)
	}
	role.Actions[roleAction.Name] =  roleAction
}