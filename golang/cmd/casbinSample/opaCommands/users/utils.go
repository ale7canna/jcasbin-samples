package users

import "golang/cmd/casbinSample/opaCommands/types"

var users = []types.User{
	Ale,
	Andrea,
	AleMS,
	Fabio,
	Luca,
}

func GetInfo() map[string]types.User {
	result := map[string]types.User{}
	for _, u := range users {
		result[u.Id] = u
	}
	return result

}
