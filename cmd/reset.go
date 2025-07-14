package cmd

import (
	"fmt"

	"github.com/AnhTTx13/askai/internal/data"
	"github.com/spf13/cobra"
)

var reset = &cobra.Command{
	Use:   "reset",
	Short: "Reset default options",
	Long: `Reset default options
Example: 
  askai reset
`,
	Run: func(cmd *cobra.Command, args []string) {
		optFile := data.NewFile[data.Options](data.AppDir, "options.json")
		opts := data.LoadOptions()
		opts.Lang = data.BasedOpts.Lang
		opts.Model = data.BasedOpts.Model
		opts.Temp = data.BasedOpts.Temp
		opts.MaxOT = data.BasedOpts.MaxOT
		opts.TopP = data.BasedOpts.TopP
		opts.TopK = data.BasedOpts.TopK
		optFile.Save(opts)
		fmt.Println("Options have been reset")
	},
}
