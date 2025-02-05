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

- ```--pro```    Use gemini-1.5-pro-latest model (default use "gemini-1.5-flash")

- ```--lang```   Specify the response language (default "English")  

- ```--help```   Help for askai  

**Example:**

```bash
askai --pro --lang Vietnamese write a story about a magic backpack.
```

**Note:**

- Never share or commit your API key
- Regularly rotate your API key for security
