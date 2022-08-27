package roles

import "golang/cmd/casbinSample/opaCommands/types"

var Anonymous = types.Role{
	Id: "anonymous",
	Policies: []types.Policy{
		{
			Id:                "anonymous-consumes-public-content",
			Action:            "consume",
			Effect:            "allow",
			FiltersOnResource: []string{"isPublic"},
			Resource:          "urn:cloudacademy:content::**",
		},
	},
}
