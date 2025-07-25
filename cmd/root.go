package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"

	"github.com/AnhTTx13/askai/internal/data"
	"github.com/AnhTTx13/askai/internal/model"
)

type ModelConfigFlag struct {
	Lang   string
	Model  string
	Temp   float32
	MaxOT  int32
	TopP   float32
	TopK   int32
	Stream bool
}

var prompt string

var (
	cf      *ModelConfigFlag
	opts    data.Options
	rootCmd *cobra.Command
)

func init() {
	cf = &ModelConfigFlag{}
	opts = data.LoadOptions()

	rootCmd = &cobra.Command{
		Use:   "askai",
		Short: "Prompt to ask ai",
		Long: `Prompt to ask ai
		
Example: 
	askai --model=gemini-2.0-flash --lang=Vietnamese --temp=2.0 --limit=4000 "write a story about a magic backpack."
	`,
		Run: func(cmd *cobra.Command, args []string) {
			if opts.ApiKey == "" {
				fmt.Println("Missing API key")
				os.Exit(1)
			}
			prompt = strings.TrimPrefix(prompt, " ")
			prompt = strings.TrimSuffix(prompt, " ")
			if prompt == "" {
				fmt.Println("Specify your prompt with -p flag")
				return
			}

			ctx := cmd.Context()
			client, _ := genai.NewClient(ctx, option.WithAPIKey(opts.ApiKey))

			model := model.NewModel(
				cf.Lang,
				cf.Model,
				*client,
				cf.Temp,
				cf.MaxOT,
				cf.TopP,
				cf.TopK,
				cf.Stream,
			)

			if model.ModelName == "gemini-1.5-flash" {
				err := model.GenAnswer(ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
				}
				return
			}

			err := model.GenAnswer(ctx, prompt)
			if err != nil {
				fmt.Println(err.Error())
				// Roll back to gemini-1.5-flash
				model.ModelName = "gemini-1.5-flash"
				err = model.GenAnswer(ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},
	}

	rootCmd.AddCommand(setApiKey)
	rootCmd.AddCommand(setOpt)
	rootCmd.AddCommand(listModels)
	rootCmd.AddCommand(listOpts)
	rootCmd.AddCommand(reset)

	rootCmd.PersistentFlags().StringVarP(&prompt, "prompt", "p", "", "Specify your prompt")
	rootCmd.PersistentFlags().
		StringVar(&cf.Lang, "lang", opts.Lang, "Specify the responses language")
	rootCmd.PersistentFlags().
		StringVar(&cf.Model, "model", opts.Model, `Specify what AI model to use
Avaiable model:
    - "gemini-2.5-pro-exp-03-25": Enhanced thinking and reasoning, multimodal understanding, advanced coding, and more
	- "gemini-2.0-flash": Next generation features, speed, and multimodal generation for a diverse variety of tasks
	- "gemini-2.0-flash-lite": A Gemini 2.0 Flash model optimized for cost efficiency and low latency
	- "gemini-1.5-flash": Fast and versatile performance across a diverse variety of tasks
	- "gemini-1.5-pro": Complex reasoning tasks requiring more intelligence
`)
	rootCmd.PersistentFlags().
		Float32Var(&cf.Temp, "temp", opts.Temp, "Controls the randomness of the output. Use higher values for more creative responses, and lower values for more deterministic responses. Values can range from [0.0, 2.0].")
	rootCmd.PersistentFlags().
		Float32Var(&cf.TopP, "topP", opts.TopP, "Changes how the model selects tokens for output. Tokens are selected from the most to least probable until the sum of their probabilities equals the topP value.")
	rootCmd.PersistentFlags().
		Int32Var(&cf.TopK, "topK", opts.TopK, "Changes how the model selects tokens for output. A topK of 1 means the selected token is the most probable among all the tokens in the model's vocabulary, while a topK of 3 means that the next token is selected from among the 3 most probable using the temperature. Tokens are further filtered based on topP with the final token selected using temperature sampling.")
	rootCmd.PersistentFlags().
		Int32Var(&cf.MaxOT, "limit", opts.MaxOT, "Sets the maximum number of tokens to include in a candidate.")
	rootCmd.PersistentFlags().
		BoolVar(&cf.Stream, "stream", false, "Enable text stream effect (like Gemini, chatGPT, etc) but can not render markdown")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
