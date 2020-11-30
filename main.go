package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ziru/cli-example/commands"
)

type app struct {
	rootCmd *cobra.Command
}

func newApp() *app {
	return &app{
		rootCmd: &cobra.Command{
			Use:   "cli",
			Short: "demo-cli",
		},
	}
}

func (a *app) addCommands(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		a.rootCmd.AddCommand(cmd)
	}
}

func main() {
	a := newApp()
	a.addCommands(
		commands.NewVersionCommand(),
		commands.NewSayCommand(),
	)
	if err := a.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
