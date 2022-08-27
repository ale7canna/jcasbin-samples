package resources

import "golang/cmd/casbinSample/opaCommands/types"

var resources = []types.Resource{
	{
		Urn:      "urn:cloudacademy:content::labs/lab123",
		IsPublic: true,
	},
	{
		Urn:       "urn:cloudacademy:content::courses/courseABC",
		IsPaywall: true,
	},
	{
		Urn:       "urn:cloudacademy:content::labs/lab-paywall",
		IsPaywall: true,
	},
}

func Get() map[string]types.Resource {
	result := map[string]types.Resource{}
	for _, r := range resources {
		result[r.Urn] = r
	}
	return result
}
