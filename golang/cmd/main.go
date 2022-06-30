package main

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/spf13/cobra"
	commands "golang/cmd/casbinSample"
	"os"
)

func main() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.InfoLevel)

	rootCmd := &cobra.Command{
		Use: "casbin-sample",
	}

	commandsManager := commands.NewManager()
	rootCmd.AddCommand(commandsManager.SetupDB())
	rootCmd.AddCommand(commandsManager.CheckPolicy())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
