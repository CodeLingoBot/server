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

type Request struct {
	User user.User
	Action string
	Resource string
}

var unconfidentDeny Result = Result{Confident:false,UserLevelAction:false, Authorized:false}

//Is a user authorized to perform an action
//User Actions are given priority over Role Actions
func IsAuthorized(request Request) Result {
	if request.Resource != `` {
		if byUserResourceActions:= isAuthorizedByUserResourceActions(request); byUserResourceActions.Confident {
			return byUserResourceActions
		}

		if byUserRoleResourceActions := isAuthorizedByRoleResourceActions(request); byUserRoleResourceActions.Confident {
			return byUserRoleResourceActions
		}
	}

	if byUserActions:= isAuthorizedByUserActions(request); byUserActions.Confident {
		return byUserActions
	}

	if byUserRoleActions := isAuthorizedByRoleActions(request); byUserRoleActions.Confident {
		return byUserRoleActions
	}
	return unconfidentDeny
}

func isAuthorizedByUserResourceActions(request Request) Result{
 	if request.User.Resources[request.Resource].Actions[request.Action].Authorized == true{
		return Result{Confident:true,UserLevelAction:true, Authorized:true}
	} else {
		return Result{Confident:true,UserLevelAction:true, Authorized:false}
	}
	return Result{Confident:false,UserLevelAction:true,Authorized:false}
}

func isAuthorizedByRoleResourceActions(request Request) Result{
	return unconfidentDeny
}

func isAuthorizedByUserActions(request Request) Result{
	for _, assignedAction := range request.User.Actions {
		if assignedAction.Name == request.Action {
			if assignedAction.Authorized == true {
				return Result{Confident:true,UserLevelAction:true, Authorized:true}
			}else{
				return Result{Confident:true,UserLevelAction:true, Authorized:false}
			}
		}
	}
	return Result{Confident:false,UserLevelAction:true,Authorized:false}
}


func isAuthorizedByRoleActions(request Request) Result{
	authorizingRoleActionExists := false
	for _,userRole := range request.User.Roles {
		if searchActions(userRole.Actions, request.Action, false) == true {
			return Result{Confident:true,UserLevelAction:false, Authorized:false}
		}

		if searchActions(userRole.Actions, request.Action, true) == true {
			authorizingRoleActionExists = true;
		}
	}
	if authorizingRoleActionExists {
		return Result{Confident:true,UserLevelAction:false, Authorized:true}
	}
	return unconfidentDeny
}

func searchActions(actions map[string]action.Action, action string, authorized bool) bool{
	for _, assignedAction := range actions {
		if assignedAction.Name == action && assignedAction.Authorized == authorized{
				return true
		}
	}
	return false;
}