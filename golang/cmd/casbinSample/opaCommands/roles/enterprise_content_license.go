package roles

import "golang/cmd/casbinSample/opaCommands/types"

var EnterpriseContentLicense = types.Role{
	Id: "enterprise-content-license",
	Policies: []types.Policy{
		{
			Id:                "enterprise-content-license-consume-company-content",
			Action:            "consume",
			Effect:            "allow",
			FiltersOnResource: []string{},
			Resource:          "urn:cloudacademy:content:{companyId}:**",
		},
		{
			Id:                "enterprise-content-license-consume-public-content",
			Action:            "consume",
			Effect:            "allow",
			FiltersOnResource: []string{},
			Resource:          "urn:cloudacademy:content::**",
		},
	},
}
