package cmd

import (
	"fmt"

	"github.com/AnhTTx13/askai/internal/data"
	"github.com/spf13/cobra"
)

var listOpts = &cobra.Command{
	Use:   "list-opts",
	Short: "List all default options",
	Long: `List all default options
Example:
  askai list-opts
`,
	Run: func(cmd *cobra.Command, args []string) {
		optFile := data.NewFile[data.Options](data.AppDir, "options.json")
		savedOpt := data.Options{}
		optFile.Load(&savedOpt)
		fmt.Printf("\nModel: %s\n", savedOpt.Model)
		fmt.Printf("Language: %s\n", savedOpt.Lang)
		fmt.Printf("Temperature: %.2f\n", savedOpt.Temp)
		fmt.Printf("MaxOutputTokens: %d\n", savedOpt.MaxOT)
		fmt.Printf("TopP: %.2f\n", savedOpt.TopP)
		fmt.Printf("TopK: %d\n", savedOpt.TopK)
	},
}
