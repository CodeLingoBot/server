package authorization_test

import (
	"testing"
	"awethome.com/authorization"
	"awethome.com/user"
	"awethome.com/action"
	"awethome.com/role"
)

func TestAllowAuthorizeOnUserAction(t *testing.T) {
	user := user.User{}
	action := action.Action{"Dance", true}
	user.AddAction(action)
	result := authorization.IsAuthorized(user, action)
	expectedResult := authorization.Result{true,true, true}
	if result != expectedResult {
		t.Error("User was not authorized on an authorized action", result)
	}
}

func TestAllowAuthorizeNotAllowedAction(t *testing.T) {
	user := user.User{}
	action := action.Action{"Dance", false}
	user.AddAction(action)
	result := authorization.IsAuthorized(user, action)
	expectedResult := authorization.Result{true,true, false}
	if result != expectedResult {
		t.Error("User was authorized on an unauthorized action", result)
	}
}

func TestDenyAuthorizeOnUnassociatedAction(t *testing.T) {
	user := user.User{}
	action := action.Action{"Dance", true}
	result := authorization.IsAuthorized(user, action)
	expectedResult := authorization.Result{false,false, false}
	if result != expectedResult {
		t.Error("User was authorized on an unassociated action", result)
	}
}

func TestAllowAuthorizeOnAllowedRoleAction(t *testing.T) {
	user := user.User{}
	role := role.Role{Name:"Dancer"}
	action := action.Action{"Dance", true}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(user, action)
	expectedResult := authorization.Result{true,false, true}
	if result != expectedResult {
		t.Error("User was not authorized on an authorized role action", result)
	}
}

func TestDenyAuthorizeOnDeniedRoleAction(t *testing.T) {
	user := user.User{}
	role := role.Role{Name:"Dancer"}
	action := action.Action{"Dance", false}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(user, action)
	expectedResult := authorization.Result{true,false, false}
	if result != expectedResult {
		t.Error("User was authorized on an unauthorized role action", result)
	}
}

func TestAllowOnRoleActionDeniesButUserActionAllows(t *testing.T) {
	user := user.User{}
	role := role.Role{Name:"Dancer"}
	allowAction := action.Action{"Dance", true}
	denyAction := action.Action{"Dance", false}
	role.AddAction(denyAction)
	user.AddRole(role)
	user.AddAction(allowAction)
	result := authorization.IsAuthorized(user, denyAction)
	expectedResult := authorization.Result{true,true, true}
	if result != expectedResult {
		t.Error("User was not authorized on user action allows but role action denies", result)
	}
}

func TestDenyByConflictingRoleActions(t *testing.T) {
	user := user.User{}
	role := role.Role{Name:"Dancer"}
	allowAction := action.Action{"Dance", true}
	denyAction := action.Action{"Dance", false}
	role.AddAction(allowAction)
	role.AddAction(denyAction)
	user.AddRole(role)
	result := authorization.IsAuthorized(user, denyAction)
	expectedResult := authorization.Result{true,false, false}
	if result != expectedResult {
		t.Error("User was authorized on conflicting role actions", result)
	}
}