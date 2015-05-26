package role

import (
	"awethome.com/action"
)

type Role struct {
	Name string
	Actions []action.Action
}

func (role *Role) AddAction(action action.Action) {
	role.Actions = append(role.Actions, action)
}
