package realm_test

import (
	"testing"
	"awethome/server/user"
	"awethome/server/role"
	"awethome/server/realm"
)

func TestAddingUsersAndRoles(t *testing.T) {
	realm := realm.Realm{}
	user := user.User{}
	realm.AddUser(user)
	if len(realm.Users) != 1 {
		t.Error(`Adding an user twice didn't work as expected`, user.Actions)
	}

	role := role.Role{Name:`dancer`}
	realm.AddRole(role)

	if realm.Roles["dancer"].Name != `dancer` {
		t.Error(`Adding a role didn't work as expected`, user.Actions)
	}
}
