package cmd

import (
	"fmt"
	"os"

	sprites "github.com/superfly/sprites-go"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:           "devpod-provider-sprites",
		Short:         "DevPod provider for Sprites.dev",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	root.AddCommand(NewInitCmd())
	root.AddCommand(NewCreateCmd())
	root.AddCommand(NewDeleteCmd())
	root.AddCommand(NewStatusCmd())
	root.AddCommand(NewCommandCmd())

	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		if exitErr, ok := err.(*sprites.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
