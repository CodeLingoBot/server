package authorization

import(
	"awethome.com/user"
	"awethome.com/action"
)

//Results of an authorization contain
//Whether there was a conclusive decision that finalizes the decision
//The action that was conclusive if any
//Whether the user is authorized or not
type Result struct {
	Final bool
	UserLevelAction bool
	Allowed bool
}

func IsAuthorized(user user.User, action action.Action) Result {
	authorizeByUserActions := isAuthorizedByUserActions(user, action)
	if authorizeByUserActions.Final == true {
		return authorizeByUserActions
	}
	return Result{false,false, false}
}

func isAuthorizedByUserActions(user user.User, action action.Action) Result{
	for _,userAction := range user.Actions {
		if userAction.Description == action.Description{
			if(userAction.Allowed == true) {
				return Result{true,true, true}
			}else{
				return Result{true,true, false}
			}
		}
	}
	return Result{false,false, false}
}