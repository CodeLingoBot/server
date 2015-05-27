package role_test

import (
	"testing"
	"awethome.com/action"
	"awethome.com/role"
)

func TestAddingActions(t *testing.T) {
	role := role.Role{}
	allowAction := action.Action{Name:"dance",Authorized:true}
	denyAction := action.Action{Name:"dance",Authorized:false}
	role.AddAction(allowAction)
	role.AddAction(denyAction)
	if len(role.Actions) != 1 {
		t.Error(`Adding an action twice didn't work as expected`, role.Actions)
	}

	if role.Actions["dance"].Authorized != false {
		t.Error(`Adding an action twice didn't update the Actions`, role.Actions)
	}
}