package users

import "golang/cmd/casbinSample/opaCommands/types"

var Andrea = types.User{
	Id:         "andrea",
	Groups:     []string{"yearly", "admin"},
	UserGrants: []types.Policy{},
}
