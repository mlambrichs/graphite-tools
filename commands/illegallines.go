package commands

import (
	"fmt"
	"github.com/mlambrichs/graphite-tools/core"
	"github.com/spf13/cobra"
	"github.com/spf13/hugo/utils"
)

// illegalLinesCmd is
var illegalLinesCmd = &cobra.Command{
	Use:   "illegallines",
	Short: "Show the illegal lines being sent to graphite",
	Long: `illegallines is a command line utility that shows all metrics that are 
    not being sent by clients in a proper way. They're being discarded and should be block by
    the relay.

    Complete documentation is available at https://github.com/mlambrichs/graphite-tools`,
	Run: func(cmd *cobra.Command, args []string) {
		value := core.IllegalLines()
	},
}

func Execute() {
	AddCommands()
	utils.StopOnErr(illegalLinesCmd.Execute())
}
