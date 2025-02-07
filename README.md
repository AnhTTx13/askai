# askai

Ask AI in terminal 😀

## Installation

**Prerequisites:**

- Go installed
- Command line interface

**1. Setup Gemini API Key:**

- Get your API key:
  - Go to [Google AI Studio](https://aistudio.google.com)  
  - Sign in with your Google account  
  - Click "Create API Key"  
  - Copy the generated key  
  
- Set environment variables (bash-Linux):  
  
  - Check if had bash's configuration file
  
    ```bash
    ~/.bashrc
    ```

  - If the response is "No such file or directory", you will need to create this file before open it

    ```bash
    touch ~/.bashrc
    open ~/.bashrc
    ```

  - Set your API key by adding the following line into ~/.bashrc file:
  
    ```bash
    export GEMINI_API_KEY=<YOUR_API_KEY>
    ```

  - After saving the file, apply the changes by running:

    ```bash
    source ~/.bashrc
    ```

  - For more about setup environment variables (Window, MacOs), read [here.](https://ai.google.dev/gemini-api/docs/api-key)

**2. Install the repository:**

```bash
go install github.com/AnhBigBrother/askai
```

## Usage

```bash
askai [flags] [your_prompt]
```

**Flags:**

- ```--lang <string>```   Specify the responses language (default "English")  

- ```--model <string>```    Specify what AI model to use (default "gemini-1.5-flash").
  Avaiable model:
  - **gemini-2.0-flash**: Next generation features, speed, and multimodal generation for a diverse variety of tasks
  - **gemini-2.0-flash-lite-preview**: A Gemini 2.0 Flash model optimized for cost efficiency and low latency
  - **gemini-1.5-flash**: Fast and versatile performance across a diverse variety of tasks
  - **gemini-1.5-pro**: Complex reasoning tasks requiring more intelligence
  _[*More about Google Gemini model.](https://ai.google.dev/gemini-api/docs/models/gemini)_

- ```--temp <float32>```   Controls the randomness of the output. Use higher values for more creative responses, and lower values for more deterministic responses. Values can range from [0.0, 2.0]. (default 1)

- ```--limit <int32>```    Sets the maximum number of tokens to include in a candidate. (default 8192)

- ```--topP <float32>```   Changes how the model selects tokens for output. Tokens are selected from the most to least probable until the sum of their probabilities equals the topP value. (default 0.95)

- ```--topK <int32>```     Changes how the model selects tokens for output. A topK of 1 means the selected token is the most probable among all the tokens in the model's vocabulary, while a topK of 3 means that the next token is selected from among the 3 most probable using the temperature. Tokens are further filtered based on topP with the final token selected using temperature sampling. (default 40)

- ```--help```   Help for askai  

**Example:**

```bash
askai --model=gemini-2.0-flash --lang=Vietnamese --temp=1.5 write a story about a magic backpack.
```

**Note:**

- Never share or commit your API key
- Regularly rotate your API key for security
