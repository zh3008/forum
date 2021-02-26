package model

//Role ..
type Role struct {
	ID   int    `json:"role_id"`
	Name string `json:"role_name"`
}

//NewRole ..
func NewRole() *Role {
	return &Role{}
}
