package models

type RealmUser struct {
	RealmId *Realm `orm:"column(realm_id);rel(fk)"`
	UserId  *User  `orm:"column(user_id);rel(fk)"`
}
