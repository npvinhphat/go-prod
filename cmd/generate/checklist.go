package generate

import (
	"github.com/spf13/cobra"
)

// checklistCmd represents the checklist command
var checklistCmd = &cobra.Command{
	Use:   "checklist",
	Short: "Checklist command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	checklistCmd.Flags().String("level", "", "Service level. Allowed: basic, intermediate, advanced.")
	checklistCmd.MarkFlagRequired("level")

	GenerateCmd.AddCommand(checklistCmd)
}
