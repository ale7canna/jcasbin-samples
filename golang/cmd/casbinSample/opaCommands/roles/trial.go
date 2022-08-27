package roles

import "golang/cmd/casbinSample/opaCommands/types"

var Trial = types.Role{
	Id: "trial",
	Policies: []types.Policy{
		{
			Id:                "trial-consume-public-content",
			Action:            "consume",
			Effect:            "allow",
			FiltersOnResource: []string{},
			Resource:          "urn:cloudacademy:content::**",
		},
		{
			Id:                "trial-cannot-consume-paywall-content",
			Action:            "consume",
			Effect:            "deny",
			FiltersOnResource: []string{"isPaywall"},
			Resource:          "urn:cloudacademy:content::**",
			DenyMessage:       "trial can't consume paywall. Action required: upsell",
		},
	},
}
