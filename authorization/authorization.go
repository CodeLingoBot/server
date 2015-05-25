package authorization

type User struct {
	FullName string
	Actions []Action
}

type Action struct {
	Description string
	Allowed bool
}

func (user *User) Authorize(action Action) {
	user.Actions = append(user.Actions, action)
}

func (user *User) IsAuthorized(action Action) bool {
	for _,userAction := range user.Actions {
		if userAction.Description == action.Description && userAction.Allowed == true{
			return true;
		}
	}
	return false
}