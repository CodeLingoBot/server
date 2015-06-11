package models

type UserRole struct {
	UserId *User `orm:"column(user_id);rel(fk)"`
	RoleId *Role `orm:"column(role_id);rel(fk)"`
}
