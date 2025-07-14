package cmd

import (
	"fmt"

	"github.com/AnhTTx13/askai/internal/data"
	"github.com/spf13/cobra"
)

var setOpt = &cobra.Command{
	Use:   "set-opt",
	Short: "Set you default app options",
	Long: `Set you default app options
Example: 
  askai set-opt --lang=Vietnamese
`,
	Run: func(cmd *cobra.Command, args []string) {
		optFile := data.NewFile[data.Options](data.AppDir, "options.json")
		opts := data.LoadOptions()
		if opts.Lang != cf.Lang {
			opts.Lang = cf.Lang
		}
		if opts.Model != cf.Model {
			opts.Model = cf.Model
		}
		if opts.Temp != cf.Temp {
			opts.Temp = cf.Temp
		}
		if opts.MaxOT != cf.MaxOT {
			opts.MaxOT = cf.MaxOT
		}
		if opts.TopP != cf.TopP {
			opts.TopP = cf.TopP
		}
		if opts.TopK != cf.TopK {
			opts.TopK = cf.TopK
		}
		optFile.Save(opts)
		fmt.Println("Your options key have been set")
	},
}
