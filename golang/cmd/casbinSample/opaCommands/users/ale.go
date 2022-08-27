package users

import "golang/cmd/casbinSample/opaCommands/types"

var Ale = types.User{
	Id:         "ale",
	Groups:     []string{"anonymous"},
	UserGrants: []types.Policy{},
}
