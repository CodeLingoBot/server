package user_test

import (
	"testing"

	"github.com/awethome/server/action"
	"github.com/awethome/server/resource"
	"github.com/awethome/server/role"
	"github.com/awethome/server/user"
)

func TestAddingActionsAndRolesResources(t *testing.T) {
	user := user.User{}
	allowAction := action.Action{Name: "dance", Authorized: true}
	denyAction := action.Action{Name: "dance", Authorized: false}
	user.AddAction(allowAction)
	user.AddAction(denyAction)
	if len(user.Actions) != 1 {
		t.Error(`Adding an action twice didn't work as expected`, user.Actions)
	}

	if user.Actions["dance"].Authorized != false {
		t.Error(`Adding an action twice didn't update the Actions`, user.Actions)
	}

	role := role.Role{Name: `dancer`}
	user.AddRole(role)
	user.AddRole(role)

	if len(user.Roles) != 1 {
		t.Error(`Adding an action twice didn't work as expected`, user.Actions)
	}

	if user.Roles["dancer"].Name != `dancer` {
		t.Error(`Adding an action twice didn't update the Actions`, user.Actions)
	}

	resource := resource.Resource{Name: `phone`}
	user.AddResource(resource)
	if user.Resources["phone"].Name != `phone` {
		t.Error(`Adding a resource twice didn't update the Actions`, user.Resources)
	}
}
