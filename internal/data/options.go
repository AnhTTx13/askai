package data

import (
	"fmt"
	"os"
)

var (
	AppDir string = ".askai"
)

type Options struct {
	Model  string  `json:"model"`
	Lang   string  `json:"language"`
	Temp   float32 `json:"temperature"`
	MaxOT  int32   `json:"max_output_token"`
	TopP   float32 `json:"top_p"`
	TopK   int32   `json:"top_k"`
	ApiKey string  `json:"api_key"`
}

var BasedOpts = Options{
	Model: "gemini-2.0-flash",
	Lang:  "English",
	Temp:  1,
	TopP:  0.95,
	TopK:  40,
	MaxOT: 8192,
}

func LoadOptions() Options {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	AppDir = fmt.Sprintf("%s/%s", homeDir, AppDir)
	optsFile := NewFile[Options](AppDir, "options.json")
	opts := Options{
		Model:  BasedOpts.Model,
		Lang:   BasedOpts.Lang,
		Temp:   BasedOpts.Temp,
		TopP:   BasedOpts.TopP,
		TopK:   BasedOpts.TopK,
		MaxOT:  BasedOpts.MaxOT,
		ApiKey: "",
	}
	err = optsFile.Load(&opts)
	if err != nil {
		optsFile.Save(opts)
	}
	if opts.ApiKey == "" {
		apiKey := os.Getenv("GEMINI_API_KEY")
		if apiKey != "" {
			opts.ApiKey = apiKey
			optsFile.Save(opts)
		}
	}
	return opts
}
