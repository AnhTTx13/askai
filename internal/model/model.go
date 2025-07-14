package model

import (
	"context"
	"fmt"

	"github.com/AnhTTx13/askai/internal/ui"
	markdown "github.com/MichaelMure/go-term-markdown"
	"golang.org/x/term"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

type Model struct {
	Lang            string
	ModelName       string
	Client          genai.Client
	Temperature     float32
	MaxOutputTokens int32
	TopP            float32
	TopK            int32
	Stream          bool
}

func NewModel(lang, model_name string, client genai.Client, temp float32, mot int32, top_p float32, top_k int32, stream bool) *Model {
	return &Model{
		Lang:            lang,
		ModelName:       model_name,
		Client:          client,
		Temperature:     temp,
		MaxOutputTokens: mot,
		TopP:            top_p,
		TopK:            top_k,
		Stream:          stream,
	}
}

func (m *Model) GenAnswer(ctx context.Context, prompt string) error {
	prompt = fmt.Sprintf("%s\nResponse in %s", prompt, m.Lang)
	fmt.Println()
	termWidth, _, err := term.GetSize(0)
	if err != nil {
		termWidth = 80
	}
	fmt.Print(string(markdown.Render(fmt.Sprintf("**Generative Model**: %s\n", m.ModelName), termWidth, 0)))

	if m.Stream {
		return m.printStream(ctx, prompt)
	}
	return m.renderMarkdown(ctx, prompt)
}

func (m *Model) printStream(ctx context.Context, prompt string) error {
	generativeModel := m.Client.GenerativeModel(m.ModelName)
	generativeModel.SetTemperature(m.Temperature)
	generativeModel.SetTopP(m.TopP)
	generativeModel.SetTopK(m.TopK)
	generativeModel.SetMaxOutputTokens(m.MaxOutputTokens)

	fmt.Println()
	iter := generativeModel.GenerateContentStream(ctx, genai.Text(prompt))
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

func (m *Model) renderMarkdown(ctx context.Context, prompt string) error {
	generativeModel := m.Client.GenerativeModel(m.ModelName)
	generativeModel.SetTemperature(m.Temperature)
	generativeModel.SetTopP(m.TopP)
	generativeModel.SetTopK(m.TopK)
	generativeModel.SetMaxOutputTokens(m.MaxOutputTokens)

	loader := ui.NewLoader([]string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[====]", "[ ===]", "[  ==]", "[   =]"}, 200)
	loader.Start()

	res, err := generativeModel.GenerateContent(ctx, genai.Text(prompt))
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
