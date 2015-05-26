package authorization

import(
	"awethome.com/user"
	"awethome.com/action"
)

//Results of an authorization contain
//Whether there was an association between this user and the action or an association between his roles and this action
//Whether the decision was made by looking at actions of the user as oposed to actions of her roles
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

	authorizeByUserRoleActions := isAuthorizedByRoleActions(user, action)
	if authorizeByUserRoleActions.Final == true {
		return authorizeByUserRoleActions
	}
	return Result{false,false, false}
}

func isAuthorizedByRoleActions(user user.User, action action.Action) Result{
	for _,userRole := range user.Roles {
		roleResult := isAuthorizedByActions(userRole.Actions, action, false)
		if roleResult.Final == true {
			return roleResult
		}
	}
	return Result{false,false, false}
}

func isAuthorizedByUserActions(user user.User, action action.Action) Result{
	return isAuthorizedByActions(user.Actions, action, true)
}

func isAuthorizedByActions(actions []action.Action, action action.Action, UserLevelAction bool) Result{
	for _, RoleAction := range actions {
		if RoleAction.Description == action.Description{
			if(RoleAction.Allowed == true) {
				return Result{true,UserLevelAction, true}
			}else{
				return Result{true,UserLevelAction, false}
			}
		}
	}
	return Result{false,UserLevelAction, false}
}