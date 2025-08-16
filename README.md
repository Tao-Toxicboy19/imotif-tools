# 🚀 IMOTIF Tools

> Generate Git commit messages interactively, quickly, and consistently.

`imotif-tools` is a cross-platform CLI tool that helps developers craft standardized commit messages via an interactive prompt or directly from the command line.

---

## ✨ Features

- 🔍 **Interactive Prompt** for Task ID, commit type, and message
- 🧠 Built-in support for common commit types (e.g. `FIX`, `ADD`, `REF`, `FEA`, `ISS`)
- 🤖 **AI-generated commit messages** with Gemini or other providers
- 🆔 Supports **multiple Task IDs** in one commit (e.g. `OD-1,OD-2,OD-3`)
- ✅ Optional commit verification (`--no-verify` support)
- 💬 Clean and readable CLI UX
- 🛠️ Works on **macOS**, **Linux**, and **Windows**
- 📦 Built-in **self-update** mechanism
- ⚡ CLI **alias support** (e.g. `itcm`)

---

## 📦 Installation

### macOS / Linux
```bash
curl -fsSL https://raw.githubusercontent.com/Tao-Toxicboy19/imotif-tools/main/install.sh | bash
```

### Windows (PowerShell)
```powershell
irm https://raw.githubusercontent.com/Tao-Toxicboy19/imotif-tools/main/install.ps1 | iex
```

---

## ⚙️ Configuration (.env)

To use the AI features, create a `.env` file in the root of your project or home directory with the following variables:

```env
GOOGLE_API_KEY=your_google_api_key
AI_PROVIDER=gemini
AI_MODEL=gemini-1.5-flash  # or any supported model
```

### Notes:
- `GOOGLE_API_KEY`: Your API key from Google AI Studio (https://makersuite.google.com/)
- `AI_PROVIDER`: The AI backend to use (`gemini`, `openai`, or `ollama`)
- `AI_MODEL`: The model name depending on provider (`gemini-1.5-flash`, `gpt-4o`, etc.)

---

## 🛠️ Setup Alias (Optional)

To shorten the CLI command, you can create an alias like:

### macOS / Linux (Zsh / Bash / Fish)
Add to your shell profile (e.g. `~/.zshrc`, `~/.bashrc`, or `~/.config/fish/functions/itcm.fish`):

```bash
alias itcm='imotif-tools commit'
alias itmg='imotif-tools magic'
```

Then reload your shell:

```bash
source ~/.zshrc  # or ~/.bashrc, depending on your shell
```

### Windows (PowerShell)

If `$PROFILE` doesn't exist:
```powershell
New-Item -ItemType File -Path $PROFILE -Force
```

Then open it:
```powershell
notepad $PROFILE
```

Add:
```powershell
function itcm {
    imotif-tools commit
}
function itmg {
    imotif-tools magic
}
```

Then reload PowerShell:
```powershell
. $PROFILE
```

---

## 🚀 Usage

### Start interactive commit:
```bash
imotif-tools commit "your message"
```

Or with alias:
```bash
itcm "your message"
```

### Auto-generate commit message with AI:
```bash
imotif-tools magic
```

> You’ll get a suggested commit message from your staged code. You can confirm or edit it.

---

## 🔄 Self-update
```bash
imotif-tools update
```

---

## 📖 Help

```bash
imotif-tools --help
```

---

## 🧠 Coming Soon

- More AI providers (OpenAI, Ollama)
- GitHub/GitLab integration
- Customizable commit templates
- Git hook support

---

Made with ❤️ by [@tao-thewarat](https://github.com/tao-thewarat)
```

หากคุณอยากเพิ่มการรองรับ config หลายระดับ (global/local), `.imotifrc` หรือ custom commit type ก็สามารถขยายจากโครงนี้ได้อีกเรื่อย ๆ ครับ ✅
