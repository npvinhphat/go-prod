package generate

import (
	"fmt"
	"os"
	"strings"

	"github.com/npvinhphat/go-prod/assets"
	"github.com/npvinhphat/go-prod/internal/data"
	"github.com/spf13/cobra"
)

const guidelinePath string = "templates/guideline_template.tmpl"

// guidelineCmd represents the guideline command
var guidelineCmd = &cobra.Command{
	Use:   "guideline",
	Short: "Generate guideline",
	Run: func(cmd *cobra.Command, args []string) {
		stacks, err := cmd.Flags().GetStringSlice("stacks")
		if err != nil {
			panic(err)
		}

		// First load the data
		var content data.Data
		content, err = data.LoadData(strings.Join(os.Args, " "), stacks)
		if err != nil {
			panic(err)
		}

		// Then apply the template
		tmpl, err := assets.Assets.ReadFile(guidelinePath)
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
	// guidelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// guidelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	GenerateCmd.AddCommand(guidelineCmd)
}
