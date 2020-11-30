package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type sayRunner struct {
	message string
}

func NewSayCommand() *cobra.Command {
	r := &sayRunner{}
	cmd := &cobra.Command{
		Use:   "say",
		Short: "Say something to me",
		Run:   r.Run,
	}
	cmd.Flags().StringVarP(&r.message, "message", "m", "", "Message to say about")
	return cmd
}

func (r *sayRunner) Run(_ *cobra.Command, _ []string) {
	fmt.Printf("Say: %s\n", r.message)
}
