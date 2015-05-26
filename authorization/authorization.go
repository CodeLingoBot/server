package authorization

import(
	"awethome.com/user"
	"awethome.com/action"
)

func IsAuthorized(user user.User, action action.Action) bool {
	for _,userAction := range user.Actions {
		if userAction.Description == action.Description{
			if(userAction.Allowed == true) {
				return true
			}else{
				return false
			}
		}
	}
	return false
}