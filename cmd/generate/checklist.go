package generate

import (
	"fmt"

	"github.com/npvinhphat/go-prod/assets"
	"github.com/npvinhphat/go-prod/internal/data"
	"github.com/spf13/cobra"
)

const checklistPath string = "templates/checklist_template.tmpl"

// checklistCmd represents the checklist command
var checklistCmd = &cobra.Command{
	Use:   "checklist",
	Short: "Generate checklist",
	Run: func(cmd *cobra.Command, args []string) {
		stacks, err := cmd.Flags().GetStringSlice("stacks")
		if err != nil {
			panic(err)
		}
		level, err := cmd.Flags().GetString("level")
		if err != nil {
			panic(err)
		}

		// First load the data
		var content map[string][]data.Item
		content, err = data.LoadData(stacks)
		if err != nil {
			panic(err)
		}

		// Then filter by level
		data.FilterData(content, level)

		// Then apply the template
		tmpl, err := assets.Assets.ReadFile(checklistPath)
		if err != nil {
			panic(err)
		}
		result, err := data.ApplyData(content, string(tmpl))
		if err != nil {
			panic(err)
		}
		fmt.Print(result)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	checklistCmd.Flags().String("level", "", "Service level. Allowed: a, b, c.")
	checklistCmd.MarkFlagRequired("level")

	GenerateCmd.AddCommand(checklistCmd)
}
