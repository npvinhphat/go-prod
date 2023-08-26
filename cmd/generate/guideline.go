package generate

import (
	"github.com/npvinhphat/go-prod/internal/data"
	"github.com/spf13/cobra"
)

// guidelineCmd represents the guideline command
var guidelineCmd = &cobra.Command{
	Use:   "guideline",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stacks, err := cmd.Flags().GetStringSlice("stacks")
		if err != nil {
			panic(err)
		}
		_, err = data.LoadData(stacks)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// guidelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guidelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	GenerateCmd.AddCommand(guidelineCmd)
}
