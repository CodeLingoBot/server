package authorization_test

import (
	"testing"
	"awethome.com/authorization"
	"awethome.com/user"
	"awethome.com/action"
	"awethome.com/role"
)

func TestAuthorizeOnAllowedAction(t *testing.T) {
	user := user.User{"John Smith", emptyActionsArray(), emptyRolesArray()}
	action := action.Action{"Dance", true}
	user.AddAction(action)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != true {
		t.Error("User was not authorized on an authorized action", result)
	}
}

func TestAuthorizeNotAllowedAction(t *testing.T) {
	user := user.User{"John Smith", emptyActionsArray(), emptyRolesArray()}
	action := action.Action{"Dance", false}
	user.AddAction(action)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != false {
		t.Error("User was authorized on an unauthorized action", result)
	}
}

func TestAuthorizeOnUnassociatedAction(t *testing.T) {
	user := user.User{"John Smith", emptyActionsArray(), emptyRolesArray()}
	action := action.Action{"Dance", true}
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != false {
		t.Error("User was authorized on an unassociated action", result)
	}
}

func TestAuthorizeOnAllowedRoleAction(t *testing.T) {
	user := user.User{"John Smith", emptyActionsArray(),  emptyRolesArray()}
	role := role.Role{"Dancer", emptyActionsArray()}
	action := action.Action{"Dance", true}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != true {
		t.Error("User was not authorized on an authorized action to a user role", result)
	}
}

func TestAuthorizeOnNotAllowedRoleAction(t *testing.T) {
	user := user.User{"John Smith", emptyActionsArray(),  emptyRolesArray()}
	role := role.Role{"Dancer", emptyActionsArray()}
	action := action.Action{"Dance", false}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(user, action)
	if result.Allowed != false {
		t.Error("User was not authorized on an unauthorized action to a user role", result)
	}
}

func emptyActionsArray() []action.Action{
	return make([]action.Action,0)
}

func emptyRolesArray() []role.Role{
	return make([]role.Role,0)
}