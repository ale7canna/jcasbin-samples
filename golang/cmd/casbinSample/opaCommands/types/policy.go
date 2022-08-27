package types

type Policy struct {
	Id                string   `json:"id"`
	Action            string   `json:"action"`
	Effect            string   `json:"effect"`
	FiltersOnResource []string `json:"filters_on_resource"`
	Resource          string   `json:"resource"`
	DenyMessage       string   `json:"deny_message"`
}
