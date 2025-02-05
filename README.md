# askai

Ask AI in terminal 😀

## Installation

**Prerequisites:**

- Go installed
- Command line interface

**1.Get the repository:**

```bash
git clone github.com/AnhBigBrother/askai
cd askai

go mod tidy
```

**2.Configure Google Gemini API:**

- Get your API key:
  - Go to [Google AI Studio](https://aistudio.google.com)  
  - Sign in with your Google account  
  - Click "Create API Key"  
  - Copy the generated key  

- Set up environment variables:  
  - Create a `.env` file in the project root  
  - Add your API key:  

     ```env  
     GOOGLE_API_KEY=your_google_api_key
     ```

  - Add `.env` to your `.gitignore` file

**3.Install binary file:**

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
- Keep your `.env` file private
- Regularly rotate your API key for security
