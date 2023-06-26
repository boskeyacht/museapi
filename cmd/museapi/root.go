package museapi

import (
	"github.com/boskeyacht/museapi/internal/types"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:  "muse api",
		Long: "muse api server",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cfg *types.Config
)

func Execute(c *types.Config) {
	cfg = c

	rootCmd.Execute()
}
