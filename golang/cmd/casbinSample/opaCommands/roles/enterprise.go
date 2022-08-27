package roles

import "golang/cmd/casbinSample/opaCommands/types"

var Enterprise = types.Role{
	Id: "enterprise",
	Policies: []types.Policy{
		{
			Id:                "enterprise-reads-users-name",
			Action:            "read-name",
			Effect:            "allow",
			FiltersOnResource: []string{},
			Resource:          "urn:cloudacademy:users:{companyId}:users/**",
		},
	},
}
