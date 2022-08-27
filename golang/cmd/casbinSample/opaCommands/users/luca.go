package users

import (
	"golang/cmd/casbinSample/opaCommands/types"
	"strconv"
)

var Luca = types.User{
	Id:     "luca",
	Groups: []string{"yearly"},
	UserGrants: append([]types.Policy{{
		Id:                "luca-consume-lab-paywall",
		Action:            "consume",
		Effect:            "allow",
		FiltersOnResource: []string{},
		Resource:          "urn:cloudacademy:content::labs/lab-paywall",
	}}, fakePolicies(1000)...),
}

func fakePolicies(n int) []types.Policy {
	result := make([]types.Policy, n)
	for i := 0; i < n; i++ {
		result[i] = types.Policy{
			Id:       "luca-policy" + strconv.Itoa(i),
			Action:   "consume",
			Effect:   "allow",
			Resource: "urn:cloudacademy:content::courses/course-" + strconv.Itoa(i),
		}
	}
	return result
}
