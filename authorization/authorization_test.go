package authorization_test

import (
	"testing"
	"awethome.com/authorization"
	"awethome.com/user"
	"awethome.com/action"
	"awethome.com/role"
)

func TestAuthorizeOnAllowedAction(t *testing.T) {
	user := user.User{"John Smith",make([]action.Action,0), make([]role.Role,0)}
	action := action.Action{"Dance", true}
	user.AddAction(action)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != true {
		t.Error("User was not authorized on an authorized action", result)
	}
}

func TestAuthorizeNotAllowedAction(t *testing.T) {
	user := user.User{"John Smith",make([]action.Action,0), make([]role.Role,0)}
	action := action.Action{"Dance", false}
	user.AddAction(action)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != false {
		t.Error("User was authorized on an unauthorized action", result)
	}
}

func TestAuthorizeOnUnassociatedAction(t *testing.T) {
	user := user.User{"John Smith",make([]action.Action,0), make([]role.Role,0)}
	action := action.Action{"Dance", true}
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != false {
		t.Error("User was authorized on an unassociated action", result)
	}
}

func TestAuthorizeOnAllowedRoleAction(t *testing.T) {
	user := user.User{"John Smith",make([]action.Action,0), make([]role.Role,0)}
	role := role.Role{"Dancer", make([]action.Action,0)}
	action := action.Action{"Dance", true}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != true {
		t.Error("User was not authorized on an authorized action to a user role", result)
	}
}
