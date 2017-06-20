package cmd

import (
	"fmt"
	"os"
	"github.com/mlambrichs/graphite-tools/core"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "illegallines",
	Short: "Show the illegal lines being sent to graphite",
	Long: `illegallines is a command line utility that shows all metrics that are
not being sent by clients in a proper way. They're being discarded and should be block by
the relay.

Complete documentation is available at https://github.com/mlambrichs/graphite-tools`,
	Run: func(cmd *cobra.Command, args []string) {
		core.IllegalLines()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize()
}
