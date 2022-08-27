package types

type Resource struct {
	Urn       string `json:"urn"`
	IsPublic  bool   `json:"isPublic"`
	IsPaywall bool   `json:"isPaywall"`
}
