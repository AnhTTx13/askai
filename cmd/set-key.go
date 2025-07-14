package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnhTTx13/askai/internal/data"
	"github.com/spf13/cobra"
)

var setApiKey = &cobra.Command{
	Use:   "set-key",
	Short: "Set you Gemini Api key",
	Long: `Set you Gemini Api key
Example: 
  askai set-key <your_api_key>
`,
	Run: func(cmd *cobra.Command, args []string) {
		api_key := strings.Join(args, "")
		if api_key == "" {
			fmt.Println("Your key is empty.")
			os.Exit(1)
		}
		optFile := data.NewFile[data.Options](data.AppDir, "options.json")
		opts := data.LoadOptions()
		opts.ApiKey = api_key
		optFile.Save(opts)
		fmt.Println("Your API key have been set")
	},
}
