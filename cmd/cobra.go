package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zok2/meow/cmd/api"
	"github.com/zok2/meow/cmd/version"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "meow",
	Short:             "meow",
	SilenceUsage:      true,
	Long:              `meow`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks 的相关内容`
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
