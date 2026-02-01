package models

type User struct {
	ID	int `json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"-"`
	RoleID	string	`json:"role_id"`

	Role	*Role	`json:"role,omitempty"`
}