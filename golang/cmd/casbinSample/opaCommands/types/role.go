package types

type Role struct {
	Id       string   `json:"id"`
	Policies []Policy `json:"policies"`
}
