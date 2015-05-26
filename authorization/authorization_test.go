package authorization_test

import (
	"testing"
	"awethome.com/authorization"
	"awethome.com/user"
	"awethome.com/action"
)

func TestAuthorizeOnAction(t *testing.T) {
	user := user.User{"John Smith",make([]action.Action,0)}
	action := action.Action{"Dance", true}
	user.Authorize(action)
	results := authorization.IsAuthorized(user, action)
	if results != true {
		t.Error("User was not authorized on an authorized action")
	}
}

func TestAuthorizeOnUnassociatedAction(t *testing.T) {
	user := user.User{"John Smith",make([]action.Action,0)}
	action := action.Action{"Dance", true}
	results := authorization.IsAuthorized(user, action)
	if results != false {
		t.Error("User was authorized on an unassociated action")
	}
}