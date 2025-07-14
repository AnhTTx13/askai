package cmd

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var listModels = &cobra.Command{
	Use:   "list-models",
	Short: "List all avaiable models",
	Long: `List all avaiable models
	
Example: 
	askai list-models
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		client, _ := genai.NewClient(ctx, option.WithAPIKey(opts.ApiKey))

		flashRegex := regexp.MustCompile(`.*gemini-[0-9]\.[0-9]-flash.*`)
		proRegex := regexp.MustCompile(`.*gemini-[0-9]\.[0-9]-pro.*`)

		iter := client.ListModels(ctx)

		fmt.Printf("\n\033[1mAvaiable models:\033[0m\n") // Bold text

		for {
			modelInfo, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}

			m_name := modelInfo.Name

			if flashRegex.MatchString(m_name) || proRegex.MatchString(m_name) {

				if !slices.Contains(modelInfo.SupportedGenerationMethods, "generateContent") {
					continue
				}

				m_name = strings.TrimPrefix(m_name, "models/")

				fmt.Println(m_name)
			}
		}
	},
}
