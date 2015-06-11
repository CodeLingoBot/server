package models

type RealmRole struct {
	RealmId *Realm `orm:"column(realm_id);rel(fk)"`
	RoleId  *Role  `orm:"column(role_id);rel(fk)"`
}
