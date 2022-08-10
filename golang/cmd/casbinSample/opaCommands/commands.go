package opaCommands

import (
	"context"
	"github.com/apex/log"
	"github.com/open-policy-agent/opa/rego"
	"github.com/spf13/cobra"
	"time"
)

type Manager struct {
}

func (m Manager) Check() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		m.check()
		return nil
	}

	cmd := &cobra.Command{
		Use:  "opa-check",
		RunE: run,
	}
	return cmd
}

func (m Manager) check() {
	log.Info("Check run")
	module := `
package opaCommands

import future.keywords

default allow := false
default is_admin := false

allow if {
    input.method == "GET"
    input.path == ["salary", input.subject.user]
}

allow if {
    is_admin
}

is_admin if {
    "admin" in input.subject.groups
}
`
	module = `
package opaCommands

import future.keywords
default allow := false
default deny := false
default result := false
default grants := []
default role_grants := []

deny if { 
	input.user == "luca"
}

result if {
	not deny
    allow
}

# Allow the action if the user is granted permission to perform the action.
allow if {
	some grant in user_is_granted
	input.action == grant.action    
	resource_matches(grant.resource)
}

# Allow the action if the user is granted permission to perform the action.
allow if {
	input.action == "consume"    
	user_is_anonymous
	resource_is_public
}

allow if {
	input.action == "consume"    
	user_is_paying
	resource_is_paywalled
}

user_is_granted contains grant if {
	some role in input.external.user_roles[input.user]
    some grant in array.concat(input.external.user_grants[input.user], input.external.role_grants[role])
#     some role_grant in 
}

resource_matches(resource_pattern) if {
	glob.match(resource_pattern, [":"], input.resource)
}

user_is_anonymous if {
	"anonymous" in input.external.user_roles[input.user]
}

user_is_paying if {
	"yearly" in input.external.user_roles[input.user]
}

resource_is_public if {
	input.external.resource_attributes[input.resource].isPublic
}

resource_is_paywalled if {
	input.external.resource_attributes[input.resource].isPaywall
}
`
	ctx := context.TODO()
	query, err := rego.New(
		rego.Query("allow = data.opaCommands.allow"),
		rego.Module("opaCommands", module),
	).PrepareForEval(ctx)

	if err != nil {
		panic(err)
	}

	input := map[string]interface{}{
		"action":   "consume",
		"resource": "content:courses:courseABC",
		"user":     "andrea",
		//"method": "GET",
		//"path":   []interface{}{"salary", "bob"},
		//"subject": map[string]interface{}{
		//	"user":   "bob",
		//	"groups": []interface{}{"sales", "marketing"},
		//},
		"external": map[string]interface{}{
			"user_roles": map[string]interface{}{
				"andrea": []interface{}{"yearly", "admin"},
				"luca":   []interface{}{"trial"},
				"ale":    []interface{}{"anonymous"},
				"fabio":  []interface{}{"anonymous"},
			},
			"role_grants": map[string]interface{}{
				"anonymous": []interface{}{},
				"admin": []interface{}{
					map[string]interface{}{
						"action":   "edit",
						"resource": "**",
					},
				},
				"trial": []interface{}{},
				"yearly": []interface{}{
					map[string]interface{}{
						"action":   "consume",
						"resource": "content:**",
					},
					map[string]interface{}{
						"action":   "read",
						"resource": "content:labs:**",
					},
				},
			},
			"user_grants": map[string]interface{}{
				"luca": []interface{}{
					map[string]interface{}{
						"action":   "consume",
						"resource": "content:courses:course123",
					},
				},
				"andrea": []interface{}{},
				"fabio": []interface{}{
					map[string]interface{}{
						"action":   "edit",
						"resource": "content:**",
					},
				},
			},
			"resource_attributes": map[string]interface{}{
				"content:labs:lab123": map[string]interface{}{
					"isPublic": true,
				},
				"content:courses:courseABC": map[string]interface{}{
					"isPaywall": true,
				},
			},
		},
	}

	start := time.Now().UnixMicro()
	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.WithError(err).Error("Error")
	} else if len(results) == 0 {
		log.Info("Undefined result") // Handle undefined result.
	} else {
		log.WithField("result", results[0].Bindings).Info("Result") // Handle undefined result.
		// Handle result/decision.
		// fmt.Printf("%+v", results) => [{Expressions:[true] Bindings:map[x:true]}]
	}
	log.WithField("timeSpent", time.Now().UnixMicro()-start).Info("Checking policy took {timeSpent} ms")
}

func NewManager() Manager {
	start := time.Now().UnixMilli()
	log.WithField("timeSpent", time.Now().UnixMilli()-start).Info("init {timeSpent}")

	return Manager{}
}
