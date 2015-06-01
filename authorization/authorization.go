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

var unconfidentDeny Result = Result{Confident: false, UserLevelAction: false, Authorized: false}

//Is a user authorized to perform an action
//User Actions are given priority over Role Actions
func IsAuthorized(request Request) Result {
	if request.Resource != `` {
		if byUserResourceActions := isAuthorizedByUserResourceActions(request); byUserResourceActions.Confident {
			return byUserResourceActions
		}

		if byUserRoleResourceActions := isAuthorizedByRoleResourceActions(request); byUserRoleResourceActions.Confident {
			return byUserRoleResourceActions
		}
	}

	if byUserActions := isAuthorizedByUserActions(request); byUserActions.Confident {
		return byUserActions
	}

	if byUserRoleActions := isAuthorizedByRoleActions(request); byUserRoleActions.Confident {
		return byUserRoleActions
	}
	return unconfidentDeny
}

//Is user authorized to perform an action on a resource based on User Resource Actions?
func isAuthorizedByUserResourceActions(request Request) Result {
	if request.User.Resources[request.Resource].Actions[request.Action].Authorized == true {
		return Result{Confident: true, UserLevelAction: true, Authorized: true}
	} else {
		if userResource, ok := request.User.Resources[request.Resource]; ok {
			if userResourceAction, ok := userResource.Actions[request.Action]; ok {
				if userResourceAction.Authorized == false {
					return Result{Confident: true, UserLevelAction: true, Authorized: false}
				}
			}
		}
	}
	return unconfidentDeny
}

//Is user authorized to perform an action on a resource based on User Role Resource Actions?
func isAuthorizedByRoleResourceActions(request Request) Result {
	authorizingRoleActionExists := false
	for _, userRole := range request.User.Roles {
		roleResource := userRole.Resources[request.Resource]
		if roleResource.Actions[request.Action].Authorized == false {
			return Result{Confident: true, UserLevelAction: false, Authorized: false}
		}

		if roleResource.Actions[request.Action].Authorized == true {
			authorizingRoleActionExists = true
		}
	}
	if authorizingRoleActionExists {
		return Result{Confident: true, UserLevelAction: false, Authorized: true}
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
