package roles

import "golang/cmd/casbinSample/opaCommands/types"

var roles = []types.Role{
	Anonymous,
	Trial,
	Yearly,
	Enterprise,
	EnterpriseContentLicense,
}

func GetRoles() map[string]types.Role {
	result := map[string]types.Role{}
	for _, r := range roles {
		result[r.Id] = r
	}
	return result
}
