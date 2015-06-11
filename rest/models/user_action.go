package models

type UserAction struct {
	UserId   *User   `orm:"column(user_id);rel(fk)"`
	ActionId *Action `orm:"column(action_id);rel(fk)"`
}
