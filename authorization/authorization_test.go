package authorization_test

import (
	"testing"
	"awethome.com/authorization"
)

func TestAuthorizeOnAction(t *testing.T) {
	user := authorization.User{"John Smith",make([]authorization.Action,0)}
	action := authorization.Action{"Dance", true}
	user.Authorize(action)
	results := user.IsAuthorized(action)
	if results != true {
		t.Error("User was not authorized on an authorized action")
	}
}


func TestAuthorizeOnUnassociatedAction(t *testing.T) {
	user := authorization.User{"John Smith",make([]authorization.Action,0)}
	action := authorization.Action{"Dance", true}
	results := user.IsAuthorized(action)
	if results != false {
		t.Error("User was authorized on an unassociated action")
	}
}