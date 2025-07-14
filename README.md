# askai

Ask AI in terminal 😀

## Installation

**Prerequisites:**

- Go installed
- Command line interface

**1. Install the repository:**

```sh
go install github.com/AnhTTx13/askai
```

**2. Setup Gemini API Key:**

- Get your API key:
  - Go to [Google AI Studio](https://aistudio.google.com)  
  - Sign in with your Google account  
  - Click "Create API Key"  
  - Copy the generated key  
  
- Set environment variables (bash/zsh) by add `export GEMINI_API_KEY=[YOUR_API_KEY]` into your bashrc/zshrc file and source it:

    ```sh
    source ~/.bashrc
    # or source ~/.zshrc
    ```

- Or you can set up your API key directly by:

    ```sh
    askai set-key [your_api_key]
    ```

## Usage

```bash
askai [command] [flags]
```

**Commands:**

- ```set-key``` Set your API key
- ```set-opt``` Set your default options
- ```list-opts``` List all your default options
- ```list-models``` List all avaiable models
- ```reset``` Reset your default options

**Flags:**

- ```--prompt (-p) [string]```   Specify your prompt  

- ```--lang [string]```   Specify the responses language (default "English")  

- ```--model [string]```    Specify what AI model to use (default "gemini-2.0-flash"). use ```askai list-models``` to list all avaiable models.

- ```--stream [boolean]``` Enable text stream effect (like Gemini, chatGPT, etc), but can not render markdown. (default false)

- ```--temp [float32]```   Controls the randomness of the output. Use higher values for more creative responses, and lower values for more deterministic responses. Values can range from [0.0, 2.0].

- ```--limit [int32]```    Sets the maximum number of tokens to include in a candidate.

- ```--topP [float32]```   Changes how the model selects tokens for output. Tokens are selected from the most to least probable until the sum of their probabilities equals the topP value.

- ```--topK [int32]```     Changes how the model selects tokens for output. A topK of 1 means the selected token is the most probable among all the tokens in the model's vocabulary, while a topK of 3 means that the next token is selected from among the 3 most probable using the temperature. Tokens are further filtered based on topP with the final token selected using temperature sampling.

- ```--help```   Help for askai  

**Example:**

```sh
askai set-opt --lang=Vietnamese --temp=1.5 --model=gemini-2.5-pro
askai -p="write a story about a magic backpack." --stream
askai -p="write a story about a magic backpack." --lang=Japanese
```

**Note:**

- Never share or commit your API key
- Regularly rotate your API key for security
