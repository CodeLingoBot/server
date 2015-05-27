package realm

import(
	"awethome.com/user"
	"awethome.com/role"
)

type Realm struct {
	Name string
	Users map[string]user.User
	Roles map[string]role.Role
}

func (realm *Realm) AddUser(realmUser user.User) {
	if len(realm.Users) == 0 {
		realm.Users = make(map[string]user.User)
	}
	realm.Users[realmUser.Id] =  realmUser
}

func (realm *Realm) AddRole(userRole role.Role) {
	if len(realm.Roles) == 0 {
		realm.Roles = make(map[string]role.Role)
	}
	realm.Roles[userRole.Name] =  userRole
}