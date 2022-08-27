package users

import "golang/cmd/casbinSample/opaCommands/types"

var AleMS = types.User{
	Id:        "ms-ale",
	Groups:    []string{"enterprise", "enterprise-content-license"},
	CompanyId: "abcde",
}
