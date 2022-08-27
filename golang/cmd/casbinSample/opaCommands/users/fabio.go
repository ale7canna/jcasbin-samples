package users

import "golang/cmd/casbinSample/opaCommands/types"

var Fabio = types.User{
	Id:     "fabio",
	Groups: []string{"anonymous"},
	UserGrants: []types.Policy{{
		Id:                "fabio-edits-content",
		Action:            "edit",
		Effect:            "allow",
		FiltersOnResource: []string{},
		Resource:          "urn:cloudacademy:content::**",
	}},
}
