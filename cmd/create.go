package cmd

import (
	"context"
	"os"

	sprites "github.com/superfly/sprites-go"
	"github.com/spf13/cobra"
	"github.com/team-alembic/devpod-sprites/pkg/options"
)

func NewCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a sprite",
		RunE: func(_ *cobra.Command, _ []string) error {
			opts, err := options.FromEnv(false)
			if err != nil {
				return err
			}
			return runCreate(context.Background(), opts)
		},
	}
}

func runCreate(ctx context.Context, opts *options.Options) error {
	client := sprites.New(opts.Token, sprites.WithDisableControl())
	defer client.Close()

	_, err := client.CreateSprite(ctx, opts.MachineID, nil)
	if err != nil {
		return err
	}

	sprite := client.Sprite(opts.MachineID)
	cmd := sprite.CommandContext(ctx, "sudo", "apt-get", "install", "-y", "uidmap")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
