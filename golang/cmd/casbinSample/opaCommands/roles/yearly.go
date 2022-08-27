package roles

import "golang/cmd/casbinSample/opaCommands/types"

var Yearly = types.Role{
	Id: "yearly",
	Policies: []types.Policy{
		{
			Id:                "yearly-consumes-all-public-content",
			Action:            "consume",
			Effect:            "allow",
			FiltersOnResource: []string{},
			Resource:          "urn:cloudacademy:content::**",
		},
	},
}
