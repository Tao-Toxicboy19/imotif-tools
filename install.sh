#!/bin/bash

set -e

BINARY_NAME="imotif-tools"
DEFAULT_INSTALL_PATH="$HOME/.local/bin"

echo "Default install path is: $DEFAULT_INSTALL_PATH"
read -p "Enter custom install path or press [Enter] to use default: " customPath

INSTALL_PATH="${customPath:-$DEFAULT_INSTALL_PATH}"

# Make sure the install path exists
mkdir -p "$INSTALL_PATH"

# Detect OS
OS=$(uname)
BINARY_URL=""

if [[ "$OS" == "Darwin" ]]; then
    BINARY_URL="https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/${BINARY_NAME}-macos"
elif [[ "$OS" == "Linux" ]]; then
    BINARY_URL="https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/${BINARY_NAME}-linux"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Download binary
curl -L "$BINARY_URL" -o "$INSTALL_PATH/$BINARY_NAME"
chmod +x "$INSTALL_PATH/$BINARY_NAME"

echo "$BINARY_NAME installed to: $INSTALL_PATH/$BINARY_NAME"
echo ""
echo "You can now run '$BINARY_NAME' from your terminal!"

# -------------------------
# Optional: Setup alias: itcm -> imotif-tools commit
# -------------------------

echo ""
read -p "Do you want to create an alias for 'imotif-tools'? [Y/n]: " addAlias
addAlias=${addAlias:-Y} # default Y

if [[ "$addAlias" =~ ^[Yy]$ ]]; then
  echo ""
  echo "Choose your shell for alias setup:"
  echo "  [1] bash"
  echo "  [2] zsh"
  echo "  [3] fish"
  read -p "Enter your shell (bash/zsh/fish): " userShell
  userShell=$(echo "$userShell" | tr '[:upper:]' '[:lower:]') # to lowercase

  echo ""
  echo "Setting up alias: itcm → imotif-tools commit"

  case "$userShell" in
    bash)
      SHELL_RC="$HOME/.bashrc"
      if ! grep -Fxq "alias itcm='imotif-tools commit'" "$SHELL_RC"; then
        echo "alias itcm='imotif-tools commit'" >> "$SHELL_RC"
        echo "Alias added to $SHELL_RC"
      else
        echo "Alias already exists in $SHELL_RC"
      fi
      ;;
    zsh)
      SHELL_RC="$HOME/.zshrc"
      if ! grep -Fxq "alias itcm='imotif-tools commit'" "$SHELL_RC"; then
        echo "alias itcm='imotif-tools commit'" >> "$SHELL_RC"
        echo "Alias added to $SHELL_RC"
      else
        echo "Alias already exists in $SHELL_RC"
      fi
      ;;
    fish)
      FISH_CONFIG="$HOME/.config/fish/config.fish"
      if ! grep -Fxq "alias itcm 'imotif-tools commit'" "$FISH_CONFIG"; then
        echo "alias itcm 'imotif-tools commit'" >> "$FISH_CONFIG"
        echo "Alias added to $FISH_CONFIG"
      else
        echo "Alias already exists in $FISH_CONFIG"
      fi
      ;;
    *)
      echo "Unknown shell: $userShell"
      echo "Please add alias manually:"
      echo "alias itcm='imotif-tools commit'"
      ;;
  esac

  echo ""
  echo "Please restart your terminal or run 'source' on your shell config to use 'itcm'"
else
  echo "Skipped creating alias."
fi
