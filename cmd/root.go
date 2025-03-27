package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	"golang.org/x/term"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
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

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func printStream(ctx context.Context, model *genai.GenerativeModel, prompt string) error {
	iter := model.GenerateContentStream(ctx, genai.Text(prompt))
	for {
		res, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		for _, p := range res.Candidates[0].Content.Parts {
			fmt.Print(p)
		}
	}
	return nil
}

func renderMarkdown(ctx context.Context, model *genai.GenerativeModel, prompt string) error {
	res, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return err
	}
	source := ""
	for _, p := range res.Candidates[0].Content.Parts {
		source += string(p.(genai.Text))
	}
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		termWidth = 80
	}
	result := markdown.Render(source, termWidth, 1)
	fmt.Println(string(result))
	return nil
}

func useModel(client *genai.Client, model_name string, ctx context.Context, prompt string) error {
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		termWidth = 80
	}
	fmt.Println(string(markdown.Render(fmt.Sprintf("**AI Model**: %s\n", model_name), termWidth, 1)))

	model := client.GenerativeModel(model_name)

	model.SetTemperature(cf.Temperature)
	model.SetTopP(cf.TopP)
	model.SetTopK(cf.TopK)
	model.SetMaxOutputTokens(cf.MaxOutputTokens)

	if cf.TextStream {
		return printStream(ctx, model, prompt)
	}
	return renderMarkdown(ctx, model, prompt)
}

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
				err := useModel(client, cf.Model, ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
				}
				return
			}

			err := useModel(client, cf.Model, ctx, prompt)
			if err != nil {
				fmt.Println(err.Error())
				err = useModel(client, "gemini-1.5-flash", ctx, prompt)
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
