package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

type CommandFlag struct {
	Lang            string
	Model           string
	Temperature     float32
	MaxOutputTokens int32
	TopP            float32
	TopK            int32
	TextStream      bool
}

var (
	cf      *CommandFlag
	rootCmd *cobra.Command
)

func init() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("missing GEMINI_API_KEY")
	}

	cf = &CommandFlag{}

	rootCmd = &cobra.Command{
		Use:   "askai",
		Short: "Prompt to ask ai",
		Long: `Prompt to ask ai
		
Example: 
	askai --model=gemini-2.0-flash --lang=Vietnamese --temp=2.0 --limit=4000 write a story about a magic backpack.
	`,
		Run: func(cmd *cobra.Command, args []string) {
			prompt := strings.Join(args, " ")
			prompt = strings.TrimPrefix(prompt, " ")
			prompt = strings.TrimSuffix(prompt, " ")

			prompt = fmt.Sprintf("%s\nResponse in %s", prompt, cf.Lang)

			ctx := cmd.Context()
			client, _ := genai.NewClient(ctx, option.WithAPIKey(apiKey))

			if cf.Model == "gemini-1.5-flash" {
				err := newModel(client, cf.Model, ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
				}
				return
			}

			err := newModel(client, cf.Model, ctx, prompt)
			if err != nil {
				fmt.Println(err.Error())
				err = newModel(client, "gemini-1.5-flash", ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&cf.Lang, "lang", "English", "Specify the responses language")
	rootCmd.PersistentFlags().StringVar(&cf.Model, "model", "gemini-1.5-flash", `Specify what AI model to use
Avaiable model: 
	- "gemini-2.0-flash": Next generation features, speed, and multimodal generation for a diverse variety of tasks
	- "gemini-2.0-flash-lite-preview": A Gemini 2.0 Flash model optimized for cost efficiency and low latency
	- "gemini-1.5-flash": Fast and versatile performance across a diverse variety of tasks
	- "gemini-1.5-pro": Complex reasoning tasks requiring more intelligence
`)
	rootCmd.PersistentFlags().Float32Var(&cf.Temperature, "temp", 1, "Controls the randomness of the output. Use higher values for more creative responses, and lower values for more deterministic responses. Values can range from [0.0, 2.0].")
	rootCmd.PersistentFlags().Float32Var(&cf.TopP, "topP", 0.95, "Changes how the model selects tokens for output. Tokens are selected from the most to least probable until the sum of their probabilities equals the topP value.")
	rootCmd.PersistentFlags().Int32Var(&cf.TopK, "topK", 40, "Changes how the model selects tokens for output. A topK of 1 means the selected token is the most probable among all the tokens in the model's vocabulary, while a topK of 3 means that the next token is selected from among the 3 most probable using the temperature. Tokens are further filtered based on topP with the final token selected using temperature sampling.")
	rootCmd.PersistentFlags().Int32Var(&cf.MaxOutputTokens, "limit", 8192, "Sets the maximum number of tokens to include in a candidate.")
	rootCmd.PersistentFlags().BoolVar(&cf.TextStream, "stream", false, "Enable text stream effect (like Gemini, chatGPT, etc) but can not render markdown")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
