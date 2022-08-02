package opaCommands

import (
	"github.com/apex/log"
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
}

func NewManager() Manager {
	start := time.Now().UnixMilli()
	log.WithField("timeSpent", time.Now().UnixMilli()-start).Info("init {timeSpent}")

	return Manager{}
}
