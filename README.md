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

## 🛠️ Setup Alias (Optional)

To shorten the CLI command, you can create an alias like:

### macOS / Linux (Zsh / Bash / Fish)
Edit your shell profile (e.g. `~/.zshrc`, `~/.bashrc`, or `~/.config/fish/functions/itcm.fish`) and add:

```bash
alias itcm='imotif-tools commit'
```

Then reload your terminal or run:

```bash
source ~/.zshrc  # or ~/.bashrc, depending on your shell
```

### Windows (PowerShell)

If you see an error like:
```
The system cannot find the path specified.
```

#### Fix:
```powershell
New-Item -ItemType File -Path $PROFILE -Force
```

#### Then open it:
```powershell
notepad $PROFILE
```

#### Add this alias function:
```powershell
function itcm {
    imotif-tools commit
}
```

Save the file, then restart PowerShell or run:

```powershell
. $PROFILE
```

---

## 🚀 Usage

### Start interactive commit:
```bash
imotif-tools commit "your message"
```

or with alias:

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

- More AI providers (e.g. OpenAI, Ollama)
- Auto-scan diffs
- GitHub/GitLab integration

---

Made with ❤️ by [@tao-thewarat](https://github.com/tao-thewarat)

หากคุณมีโครงสร้าง project หรือ feature เพิ่มเติมในอนาคต เช่น custom config, git hooks, หรือ AI models เสริม ก็สามารถเพิ่ม section เพิ่มใน README ได้ต่อเลยครับ
