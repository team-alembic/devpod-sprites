package cmd

import (
	"context"

	sprites "github.com/superfly/sprites-go"
	"github.com/spf13/cobra"
	"github.com/team-alembic/devpod-sprites/pkg/options"
)

func NewDeleteCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a sprite",
		RunE: func(_ *cobra.Command, _ []string) error {
			opts, err := options.FromEnv(false)
			if err != nil {
				return err
			}
			return runDelete(context.Background(), opts)
		},
	}
}

func runDelete(ctx context.Context, opts *options.Options) error {
	client := sprites.New(opts.Token, sprites.WithDisableControl())
	defer client.Close()

	return client.DeleteSprite(ctx, opts.MachineID)
}
