# askai

Ask AI in terminal 😀

## Installation

**Prerequisites:**

- Go installed
- Command line interface

**1. Get the repository:**

```bash
git clone github.com/AnhBigBrother/askai
cd askai

go mod tidy
```

**2. Configure Google Gemini API:**

- Get your API key:
  - Go to [Google AI Studio](https://aistudio.google.com)  
  - Sign in with your Google account  
  - Click "Create API Key"  
  - Copy the generated key  
  
- Set up environment variables (bash-Linux):  
  
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

**3. Install binary file:**

```bash
go install
```

## Usage

```bash
askai [flags] [your_promt]
```

**Flags:**

- ```--lang```,  ```-l```   Specify the response language (default "English")  

- ```--help```, ```-h```   Help for askai  

**Example:**

```bash
askai --lang Vietnamese write a story about a magic backpack.
```

**Note:**

- Never share or commit your API key
- Regularly rotate your API key for security
