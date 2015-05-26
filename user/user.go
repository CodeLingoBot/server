package user

import (
	"awethome.com/action"
)

type User struct {
	FullName string
	Actions []action.Action
}

func (user *User) Authorize(action action.Action) {
	user.Actions = append(user.Actions, action)
}