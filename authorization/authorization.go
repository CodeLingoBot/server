package authorization

import (
	"github.com/awethome/server/action"
	"github.com/awethome/server/user"
)

//Results of an authorization contain
//Whether action was associated with user in any way or not
//Whether the decision was made by looking at actions of the user as opposed to actions of her roles
//Whether the user is authorized or not
type Result struct {
	Confident       bool
	UserLevelAction bool
	Authorized      bool
}

type Request struct {
	User     user.User
	Action   string
	Resource string
}

//An unconfident Deny occurs when we deny not because the user is really denied but because we can not
//Find a reason to allow the user, and we default to deny
var unconfidentDeny Result = Result{Confident: false, UserLevelAction: false, Authorized: false}

//Is a user authorized to perform an action
//User Actions are given priority over Role Actions
func IsAuthorized(request Request) Result {
	if byUserActions := isAuthorizedByUserActions(request); byUserActions.Confident {
		return byUserActions
	}

	if byUserRoleActions := isAuthorizedByRoleActions(request); byUserRoleActions.Confident {
		return byUserRoleActions
	}
	return unconfidentDeny
}

//Is user authorized to perform an action based on User Actions?
func isAuthorizedByUserActions(request Request) Result {
	if request.User.Actions[request.Action].Authorized == true {
		return Result{Confident: true, UserLevelAction: true, Authorized: true}
	} else {
		if userAction, ok := request.User.Actions[request.Action]; ok {
			if userAction.Authorized == false {
				return Result{Confident: true, UserLevelAction: true, Authorized: false}
			}
		}
	}
	return unconfidentDeny
}

//Is user authorized to perform an action based on User Role Actions?
func isAuthorizedByRoleActions(request Request) Result {
	authorizingRoleActionExists := false
	for _, userRole := range request.User.Roles {
		if searchActions(userRole.Actions, request.Action, false) == true {
			return Result{Confident: true, UserLevelAction: false, Authorized: false}
		}

		if searchActions(userRole.Actions, request.Action, true) == true {
			authorizingRoleActionExists = true
		}
	}
	if authorizingRoleActionExists {
		return Result{Confident: true, UserLevelAction: false, Authorized: true}
	}
	return unconfidentDeny
}

func searchActions(actions map[string]action.Action, action string, authorized bool) bool {

	if userAction, ok := actions[action]; ok {
		if userAction.Authorized == authorized {
			return true
		}
	}

	return false
}
