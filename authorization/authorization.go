package authorization

import (
	"awethome/server/user"
	"awethome/server/action"
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

var unconfidentDeny Result = Result{Confident:false, UserLevelAction:false, Authorized:false}

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

func isAuthorizedByUserResourceActions(request Request) Result {
	if request.User.Resources[request.Resource].Actions[request.Action].Authorized == true {
		return Result{Confident:true, UserLevelAction:true, Authorized:true}
	} else {
		if userResource, ok := request.User.Resources[request.Resource]; ok {
			if userResourceAction, ok := userResource.Actions[request.Action]; ok {
				if userResourceAction.Authorized == false {
					return Result{Confident:true, UserLevelAction:true, Authorized:false}
				}
			}
		}
	}
	return unconfidentDeny
}

func isAuthorizedByRoleResourceActions(request Request) Result {
	authorizingRoleActionExists := false
	for _, userRole := range request.User.Roles {
		roleResource := userRole.Resources[request.Resource];
		if roleResource.Actions[request.Action].Authorized == false {
			return Result{Confident:true, UserLevelAction:false, Authorized:false}
		}

		if roleResource.Actions[request.Action].Authorized == true {
			authorizingRoleActionExists = true;
		}
	}
	if authorizingRoleActionExists {
		return Result{Confident:true, UserLevelAction:false, Authorized:true}
	}
	return unconfidentDeny
}

func isAuthorizedByUserActions(request Request) Result {
	for _, assignedAction := range request.User.Actions {
		if assignedAction.Name == request.Action {
			if assignedAction.Authorized == true {
				return Result{Confident:true, UserLevelAction:true, Authorized:true}
			}else {
				return Result{Confident:true, UserLevelAction:true, Authorized:false}
			}
		}
	}
	return unconfidentDeny
}


func isAuthorizedByRoleActions(request Request) Result {
	authorizingRoleActionExists := false
	for _, userRole := range request.User.Roles {
		if searchActions(userRole.Actions, request.Action, false) == true {
			return Result{Confident:true, UserLevelAction:false, Authorized:false}
		}

		if searchActions(userRole.Actions, request.Action, true) == true {
			authorizingRoleActionExists = true;
		}
	}
	if authorizingRoleActionExists {
		return Result{Confident:true, UserLevelAction:false, Authorized:true}
	}
	return unconfidentDeny
}

func searchActions(actions map[string]action.Action, action string, authorized bool) bool {
	for _, assignedAction := range actions {
		if assignedAction.Name == action && assignedAction.Authorized == authorized {
			return true
		}
	}
	return false;
}