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

func IsAuthorized(user user.User, action action.Action) Result {
	if byUserActions:= isAuthorizedByUserLevelActions(user, action); byUserActions.Confident {
		return byUserActions
	}

	if byUserRoleActions := isAuthorizedByRoleActions(user, action); byUserRoleActions.Confident {
		return byUserRoleActions
	}
	return unconfidentDeny
}

func isAuthorizedByRoleActions(user user.User, action action.Action) Result{

	existsAuthorizingRole := false
	for _,userRole := range user.Roles {
		if isActionDeniedByRoleActions(userRole.Actions, action) == true {
			return Result{Confident:true,UserLevelAction:false, Authorized:false}
		}

		if isActionAllowedByRoleActions(userRole.Actions, action) == true {
			existsAuthorizingRole = true;
		}
	}
	if existsAuthorizingRole {
		return Result{Confident:true,UserLevelAction:false, Authorized:true}
	}
	return unconfidentDeny
}

func isActionDeniedByRoleActions(actions []action.Action, action action.Action) bool{
	for _, assignedAction := range actions {
		if assignedAction.Description == action.Description && assignedAction.Authorized == false{
				return true
		}
	}
	return false;
}

func isActionAllowedByRoleActions(actions []action.Action, action action.Action) bool{
	for _, assignedAction := range actions {
		if assignedAction.Description == action.Description && assignedAction.Authorized == true{
			return true
		}
	}
	return false;
}

func isAuthorizedByUserLevelActions(user user.User, action action.Action) Result{
	return isAuthorizedByActions(user.Actions, action, true)
}

func isAuthorizedByActions(actions []action.Action, action action.Action, UserLevelAction bool) Result{
	for _, assignedAction := range actions {
		if assignedAction.Description == action.Description{
			if(assignedAction.Authorized == true) {
				return Result{Confident:true,UserLevelAction:UserLevelAction, Authorized:true}
			}else{
				return Result{Confident:true,UserLevelAction:UserLevelAction, Authorized:false}
			}
		}
	}
	return Result{Confident:false,UserLevelAction:UserLevelAction,Authorized:false}
}