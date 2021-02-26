package model

//Subforum ..
type Subforum struct {
	ID   int    `json:"subforum_id"`
	Name string `json:"subforum_name"`
}

//NewSubforum ..
func NewSubforum() *Subforum {
	return &Subforum{}
}
