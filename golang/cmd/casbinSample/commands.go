package commands

import (
	"github.com/apex/log"
	"github.com/spf13/cobra"
)

type CommandsManager struct {
}

func NewManager() CommandsManager {
	return CommandsManager{}
}

func (m CommandsManager) FakeCommand() *cobra.Command {
	run := func(cmd *cobra.Command, args []string) error {
		log.Info("Command completed")
		return nil
	}

	cmd := &cobra.Command{
		Use:  "fake",
		RunE: run,
	}
	return cmd
}
