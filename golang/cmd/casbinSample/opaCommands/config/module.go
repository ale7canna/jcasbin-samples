package config

func GetModule() string {
	return `
package opaCommands

import future.keywords

default _authz_ := false
default allow := false
default grants := []
default role_grants := []

_authz_ if {
	deny == set()
	allow
}

conditions contains condition if {
	some grant in user_grants
	input.action == grant.action
	resource_matches(grant.resource)
	condition := {"id": grant.id, "filters": grant.filters_on_resource}
}

grant_condition_matches(g) if {
	every filter in g.filters_on_resource {
		resource_has_attribute(filter)
	}
}

allow if {
	some grant in user_grants
    grant.effect == "allow"
	input.action == grant.action
	resource_matches(grant.resource)
	grant_condition_matches(grant)
}


deny[msg] {
	some grant in user_grants
	grant.effect == "deny"
	input.action == grant.action
	resource_matches(grant.resource)
	grant_condition_matches(grant)
    msg := grant.deny_message
}

user_grants contains grant if {
	some role in input.external.user_info[input.user].groups
	some grant in array.concat(input.external.user_info[input.user].user_grants, input.external.role_grants[role].policies)
}

resource_matches(resource_pattern) if {
	resource := replace(resource_pattern, "{companyId}", input.external.user_info[input.user].companyId)
	glob.match(resource, [":"], input.resource)
}

resource_has_attribute(attribute) if {
	input.external.resource_attributes[input.resource][attribute] == true
}
`
}
