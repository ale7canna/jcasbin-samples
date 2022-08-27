package opaCommands

import (
	"context"
	"github.com/apex/log"
	"github.com/open-policy-agent/opa/rego"
	"github.com/spf13/cobra"
	"golang/cmd/casbinSample/opaCommands/config"
	"golang/cmd/casbinSample/opaCommands/resources"
	"golang/cmd/casbinSample/opaCommands/roles"
	"golang/cmd/casbinSample/opaCommands/users"
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
	module := config.GetModule()
	ctx := context.TODO()
	query, err := rego.New(
		rego.Query(
			"allow = data.opaCommands.allow;"+
				"_authz_ = data.opaCommands._authz_;"+
				"deny = data.opaCommands.deny;"+
				"user_grants = data.opaCommands.user_grants"),
		rego.Module("opaCommands", module),
	).PrepareForEval(ctx)

	if err != nil {
		panic(err)
	}

	input := map[string]interface{}{
		"action":   "consume",
		"resource": "urn:cloudacademy:content::labs/lab-paywall",
		"user":     "luca",
		"external": map[string]interface{}{
			"role_grants":         roles.GetRoles(),
			"user_info":           users.GetInfo(),
			"resource_attributes": resources.Get(),
		},
	}

	start := time.Now().UnixMicro()
	results, err := query.Eval(ctx, rego.EvalInput(input))
	log.WithField("timeSpent", time.Now().UnixMicro()-start).Info("Checking policy took {timeSpent} us")
	start = time.Now().UnixMicro()
	if err != nil {
		log.WithError(err).Error("Error")
	} else if len(results) == 0 {
		log.Info("Undefined result") // Handle undefined result.
	} else {
		log.WithField("authz", results[0].Bindings["_authz_"]).Info("Authz result:") // Handle undefined result.
		log.WithField("allow", results[0].Bindings["allow"]).Info("Allow:")          // Handle undefined result.
		log.WithField("deny", results[0].Bindings["deny"]).Info("Deny:")             // Handle undefined result.
		//log.WithField("user_grants", results[0].Bindings["user_grants"]).Info("User grants:") // Handle undefined result.
	}
	log.WithField("timeSpent", time.Now().UnixMicro()-start).Info("Checking policy took {timeSpent} us")

	for i := 0; i < 10; i++ {
		start := time.Now().UnixMicro()
		results, err = query.Eval(ctx, rego.EvalInput(input))
		log.WithField("timeSpent", time.Now().UnixMicro()-start).Info("Checking policy took {timeSpent} us")
	}
}

func NewManager() Manager {
	start := time.Now().UnixMilli()
	log.WithField("timeSpent", time.Now().UnixMilli()-start).Info("init {timeSpent}")

	return Manager{}
}
