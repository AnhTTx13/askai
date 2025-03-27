package cmd

import (
	"context"
	"fmt"

	termloader "github.com/AnhBigBrother/askai/pkgs/term-loader"
	markdown "github.com/MichaelMure/go-term-markdown"
	"golang.org/x/term"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

func newModel(client *genai.Client, model_name string, ctx context.Context, prompt string) error {
	fmt.Println()
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		termWidth = 80
	}
	fmt.Print(string(markdown.Render(fmt.Sprintf("**AI Model**: %s\n", model_name), termWidth, 0)))

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

func printStream(ctx context.Context, model *genai.GenerativeModel, prompt string) error {
	fmt.Println()
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
	loader := termloader.NewLoader([]string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[====]", "[ ===]", "[  ==]", "[   =]"}, 200)
	loader.Start()
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
	result := markdown.Render(source, termWidth, 0)
	loader.Stop()
	fmt.Println()
	fmt.Println(string(result))
	return nil
}
