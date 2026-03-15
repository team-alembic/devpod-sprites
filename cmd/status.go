package cmd

import (
	"context"
	"fmt"
	"net/http"

	sprites "github.com/superfly/sprites-go"
	"github.com/spf13/cobra"
	"github.com/team-alembic/devpod-sprites/pkg/options"
)

func NewStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get sprite status",
		RunE: func(_ *cobra.Command, _ []string) error {
			opts, err := options.FromEnv(false)
			if err != nil {
				return err
			}
			return runStatus(context.Background(), opts)
		},
	}
}

func runStatus(ctx context.Context, opts *options.Options) error {
	client := sprites.New(opts.Token, sprites.WithDisableControl())
	defer client.Close()

	_, err := client.GetSprite(ctx, opts.MachineID)
	if err != nil {
		if apiErr := sprites.IsAPIError(err); apiErr != nil && apiErr.StatusCode == http.StatusNotFound {
			fmt.Println("NotFound")
			return nil
		}
		return err
	}

	fmt.Println("Running")
	return nil
}
