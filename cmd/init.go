package cmd

import (
	"context"

	sprites "github.com/superfly/sprites-go"
	"github.com/spf13/cobra"
	"github.com/team-alembic/devpod-sprites/pkg/options"
)

func NewInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Validate provider configuration",
		RunE: func(_ *cobra.Command, _ []string) error {
			opts, err := options.FromEnv(true)
			if err != nil {
				return err
			}
			return runInit(context.Background(), opts)
		},
	}
}

func runInit(ctx context.Context, opts *options.Options) error {
	client := sprites.New(opts.Token, sprites.WithDisableControl())
	defer client.Close()

	_, err := client.ListSprites(ctx, nil)
	return err
}
