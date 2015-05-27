package user_test

import (
	"testing"
	"awethome.com/user"
	"awethome.com/action"
	"awethome.com/role"
)

func TestAddingActionsAndRoles(t *testing.T) {
	user := user.User{}
	allowAction := action.Action{Name:"dance",Authorized:true}
	denyAction := action.Action{Name:"dance",Authorized:false}
	user.AddAction(allowAction)
	user.AddAction(denyAction)
	if len(user.Actions) != 1 {
		t.Error(`Adding an action twice didn't work as expected`, user.Actions)
	}

	if user.Actions["dance"].Authorized != false {
		t.Error(`Adding an action twice didn't update the Actions`, user.Actions)
	}

	role := role.Role{Name:`dancer`}
	user.AddRole(role)
	user.AddRole(role)

	if len(user.Actions) != 1 {
		t.Error(`Adding an action twice didn't work as expected`, user.Actions)
	}

	if user.Actions["dance"].Authorized != false {
		t.Error(`Adding an action twice didn't update the Actions`, user.Actions)
	}
}