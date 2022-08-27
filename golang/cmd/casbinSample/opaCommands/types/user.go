package types

type User struct {
	Id         string   `json:"id"`
	CompanyId  string   `json:"companyId"`
	Groups     []string `json:"groups"`
	UserGrants []Policy `json:"user_grants"`
}
