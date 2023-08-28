package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var allowedStacks = []string{"default", "kubernetes"}

var stacks []string

// generateCmd represents the generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate docs",
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
	GenerateCmd.PersistentFlags().StringSliceVar(&stacks, "stacks", []string{"default"}, "technology stacks used")

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
