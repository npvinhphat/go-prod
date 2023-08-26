package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var allowedStacks = []string{"general", "kubernetes"}

var stacks []string

// generateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return validateStacks(stacks)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	GenerateCmd.PersistentFlags().StringSliceVar(&stacks, "stacks", []string{"general"}, "technology stacks used")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func validateStacks(stacks []string) error {
	for _, stack := range stacks {
		if !contains(allowedStacks, stack) {
			return fmt.Errorf("stack %s is not allowed", stack)
		}
	}
	return nil
}

func contains(list []string, val string) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}
