package role_test

import (
	"testing"
	"awethome.com/action"
	"awethome.com/role"
	"awethome.com/resource"
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

	resource := resource.Resource{Name:`phone`}
	role.AddResource(resource)
	if role.Resources["phone"].Name != `phone` {
		t.Error(`Adding a resource twice didn't update the Actions`, role.Resources)
	}
}