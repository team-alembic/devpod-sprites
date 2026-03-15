package cmd

import (
	"context"
	"os"

	sprites "github.com/superfly/sprites-go"
	"github.com/spf13/cobra"
	"github.com/team-alembic/devpod-sprites/pkg/options"
)

func NewCommandCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "command",
		Short: "Execute a command on the sprite",
		RunE: func(_ *cobra.Command, _ []string) error {
			opts, err := options.FromEnv(false)
			if err != nil {
				return err
			}
			return runCommand(context.Background(), opts)
		},
	}
}

func runCommand(ctx context.Context, opts *options.Options) error {
	command := os.Getenv("COMMAND")

	client := sprites.New(opts.Token, sprites.WithDisableControl())
	defer client.Close()

	sprite := client.Sprite(opts.MachineID)
	cmd := sprite.CommandContext(ctx, "sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
