package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	commands "golang/cmd/casbinSample"
	"golang/cmd/casbinSample/opaCommands"
	"os"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.InfoLevel)

	rootCmd := &cobra.Command{
		Use: "casbin-sample",
	}

	commandsManager := commands.NewManager()
	opaCommandsManager := opaCommands.NewManager()

	rootCmd.AddCommand(commandsManager.SetupDB())
	rootCmd.AddCommand(commandsManager.CheckPolicy())
	rootCmd.AddCommand(commandsManager.Benchmark())
	rootCmd.AddCommand(commandsManager.Interactive())

	rootCmd.AddCommand(opaCommandsManager.Check())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
