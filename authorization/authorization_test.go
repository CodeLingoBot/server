package authorization_test

import (
	"testing"

	"github.com/awethome/server/action"
	"github.com/awethome/server/authorization"
	"github.com/awethome/server/role"
	"github.com/awethome/server/user"
)

func TestAllowAuthorizeOnUserAction(t *testing.T) {
	user := user.User{}
	action := action.Action{Name: `dance`, Authorized: true}
	user.AddAction(action)
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{true, true, true}
	if result != expectedResult {
		t.Error(`User was not authorized on an authorized action`, result)
	}
}

func TestAllowAuthorizeNotAllowedAction(t *testing.T) {
	user := user.User{}
	action := action.Action{Name: `dance`, Authorized: false}
	user.AddAction(action)
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{true, true, false}
	if result != expectedResult {
		t.Error(`User was authorized on an unauthorized action`, result)
	}
}

func TestDenyAuthorizeOnUnassociatedAction(t *testing.T) {
	user := user.User{}
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{false, false, false}
	if result != expectedResult {
		t.Error(`User was authorized on an unassociated action`, result)
	}
}

func TestAllowAuthorizeOnAllowedRoleAction(t *testing.T) {
	user := user.User{}
	role := role.Role{Name: `dancer`}
	action := action.Action{Name: `dance`, Authorized: true}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{true, false, true}
	if result != expectedResult {
		t.Error(`User was not authorized on an authorized role action`, result)
	}
}

func TestDenyAuthorizeOnDeniedRoleAction(t *testing.T) {
	user := user.User{}
	role := role.Role{Name: `dancer`}
	action := action.Action{Name: `dance`, Authorized: false}
	role.AddAction(action)
	user.AddRole(role)
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{true, false, false}
	if result != expectedResult {
		t.Error(`User was authorized on an unauthorized role action`, result)
	}
}

func TestAllowOnRoleActionDeniesButUserActionAllows(t *testing.T) {
	user := user.User{}
	role := role.Role{Name: `dancer`}
	allowAction := action.Action{Name: `dance`, Authorized: true}
	denyAction := action.Action{Name: `dance`, Authorized: false}
	role.AddAction(denyAction)
	user.AddRole(role)
	user.AddAction(allowAction)
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{true, true, true}
	if result != expectedResult {
		t.Error(`User was not authorized on user action allows but role action denies`, result)
	}
}

func TestDenyByConflictingRoleActions(t *testing.T) {
	user := user.User{}
	role1 := role.Role{Name: `dancer`}
	role2 := role.Role{Name: `Not dancer`}
	allowAction := action.Action{Name: `dance`, Authorized: true}
	denyAction := action.Action{Name: `dance`, Authorized: false}
	role1.AddAction(allowAction)
	role2.AddAction(denyAction)
	user.AddRole(role1)
	user.AddRole(role2)
	result := authorization.IsAuthorized(authorization.Request{User: user, Action: `dance`})
	expectedResult := authorization.Result{true, false, false}
	if result != expectedResult {
		t.Error(`User was authorized on conflicting role actions`, result)
	}
}

func TestTruthTableValuesForRolesAndUsers(t *testing.T) {
	truthTable := make([][]bool, 4)
	truthTable[0] = []bool{true, true, true}
	truthTable[1] = []bool{true, false, true}
	truthTable[2] = []bool{false, true, false}
	truthTable[3] = []bool{false, false, false}

	for i := range truthTable {
		row := truthTable[i]
		userAuthorized := row[0]
		roleAuthorized := row[1]

		action1 := action.Action{Name: `action`, Authorized: userAuthorized}
		user := user.User{}
		user.AddAction(action1)

		action2 := action.Action{Name: `action`, Authorized: roleAuthorized}
		role := role.Role{}
		role.AddAction(action2)

		result := authorization.IsAuthorized(authorization.Request{User: user, Action: `action`})
		expectedResult := row[2]
		if result.Authorized != expectedResult {
			t.Error(`Expected `, expectedResult, ` in truth table (rows and users) row `, i, ` but got `, result)
		}
	}

}
