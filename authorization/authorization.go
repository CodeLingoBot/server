package authorization

import(
	"awethome.com/user"
	"awethome.com/action"
)

//Results of an authorization contain
//Whether action was associated with user in any way or not
//Whether the decision was made by looking at actions of the user as opposed to actions of her roles
//Whether the user is authorized or not
type Result struct {
	Confident bool
	UserLevelAction bool
	Authorized bool
}

var unconfidentDeny Result = Result{Confident:false,UserLevelAction:false, Authorized:false}

func IsAuthorized(user user.User, action string) Result {
	if byUserActions:= isAuthorizedByUserActions(user, action); byUserActions.Confident {
		return byUserActions
	}

	if byUserRoleActions := isAuthorizedByRoleActions(user, action); byUserRoleActions.Confident {
		return byUserRoleActions
	}
	return unconfidentDeny
}

func isAuthorizedByUserActions(user user.User, action string) Result{
	for _, assignedAction := range user.Actions {
		if assignedAction.Name == action {
			if(assignedAction.Authorized == true) {
				return Result{Confident:true,UserLevelAction:true, Authorized:true}
			}else{
				return Result{Confident:true,UserLevelAction:true, Authorized:false}
			}
		}
	}
	return Result{Confident:false,UserLevelAction:true,Authorized:false}
}

func isAuthorizedByRoleActions(user user.User,action string) Result{
	authorizingRoleActionExists := false
	for _,userRole := range user.Roles {
		if searchActions(userRole.Actions, action, false) == true {
			return Result{Confident:true,UserLevelAction:false, Authorized:false}
		}

		if searchActions(userRole.Actions, action, true) == true {
			authorizingRoleActionExists = true;
		}
	}
	if authorizingRoleActionExists {
		return Result{Confident:true,UserLevelAction:false, Authorized:true}
	}
	return unconfidentDeny
}

func searchActions(actions []action.Action, action string, expectedAuthorization bool) bool{
	for _, assignedAction := range actions {
		if assignedAction.Name == action && assignedAction.Authorized == expectedAuthorization{
				return true
		}
	}
	return false;
}