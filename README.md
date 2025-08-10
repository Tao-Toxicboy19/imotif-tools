# 🚀 IMOTIF Tools

> Generate Git commit messages interactively, quickly, and consistently.

`imotif-tools` is a cross-platform CLI tool that helps developers craft standardized commit messages via an interactive prompt or directly from the command line.

---

## ✨ Features

- 🔍 **Interactive Prompt** for Task ID, commit type, and message
- 🧠 Built-in support for common commit types (e.g. `FIX`, `ADD`, `REF`, `FEA`, `ISS`)
- 🆔 Supports **multiple Task IDs** in one commit (e.g. `OD-1,OD-2,OD-3`)
- ✅ Optional commit verification (`--no-verify` support)
- 🛠️ Works on **macOS**, **Linux**, and **Windows**
- 📦 Self-update command to get the latest version
- 💬 CLI output with clear guidance
- ⚡ Alias setup (`itcm`) for quick access

---

## 📦 Installation

### macOS / Linux
```bash
curl -fsSL https://raw.githubusercontent.com/Tao-Toxicboy19/imotif-tools/main/install.sh | bash
```

### Windows
```powershell
irm https://raw.githubusercontent.com/Tao-Toxicboy19/imotif-tools/main/install.ps1 | iex
```

### If you see
The system cannot find the path specified.

### Fix:
```powershell
New-Item -ItemType File -Path $PROFILE -Force
```

### Open it in Notepad
```
notepad $PROFILE
```

### Add the alias function inside the file:
```
function itcm {
    imotif-tools commit
}
```

### Save & restart PowerShell, or reload the profile:
