package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type CommandFlag struct {
	Lang       string
	IsProModel bool
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

func printResponse(res *genai.GenerateContentResponse) {
	for _, p := range res.Candidates[0].Content.Parts {
		fmt.Print(p)
	}
}

func useModel(client *genai.Client, model_name string, ctx context.Context, prompt string) error {
	fmt.Printf("\nused model: %s\n\n", model_name)

	model := client.GenerativeModel(model_name)

	iter := model.GenerateContentStream(ctx, genai.Text(prompt))
	for {
		res, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		printResponse(res)
	}
	return nil
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
	askai --pro --lang Vietnamese write a story about a magic backpack.
	`,
		Run: func(cmd *cobra.Command, args []string) {
			prompt := strings.Join(args, " ")
			prompt = strings.TrimPrefix(prompt, " ")
			prompt = strings.TrimSuffix(prompt, " ")
			prompt = strings.TrimSuffix(prompt, ".")

			prompt = fmt.Sprintf("%s. Response in %s", prompt, cf.Lang)

			ctx := cmd.Context()
			client, _ := genai.NewClient(ctx, option.WithAPIKey(apiKey))

			if cf.IsProModel {
				err := useModel(client, "gemini-1.5-pro-latest", ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
					err = useModel(client, "gemini-1.5-flash", ctx, prompt)
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			} else {
				err := useModel(client, "gemini-1.5-flash", ctx, prompt)
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		},
	}

	rootCmd.PersistentFlags().StringVar(&cf.Lang, "lang", "English", "Specify the responses language")
	rootCmd.PersistentFlags().BoolVar(&cf.IsProModel, "pro", false, `Use gemini-1.5-pro-latest model (default use "gemini-1.5-flash")`)
}
