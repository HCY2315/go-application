package cmd

import (
	"errors"
	"os"

	"go-application/cmd/api"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "go-application",
	Short:        "go-application",
	SilenceUsage: true,
	Long:         "go-application",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("aaa")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {

}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
