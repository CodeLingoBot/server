package models

type RoleAction struct {
	RoleId   *Role   `orm:"column(role_id);rel(fk)"`
	ActionId *Action `orm:"column(action_id);rel(fk)"`
}
