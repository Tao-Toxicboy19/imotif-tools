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
# Setup alias: itcm -> imotif-tools commit
# -------------------------

# Detect shell
SHELL_NAME=$(basename "$SHELL")

echo ""
echo "🔗 Setting up alias: itcm → imotif-tools commit"

case "$SHELL_NAME" in
  bash)
    SHELL_RC="$HOME/.bashrc"
    echo "alias itcm='imotif-tools commit'" >> "$SHELL_RC"
    echo "Alias added to $SHELL_RC"
    ;;
  zsh)
    SHELL_RC="$HOME/.zshrc"
    echo "alias itcm='imotif-tools commit'" >> "$SHELL_RC"
    echo "Alias added to $SHELL_RC"
    ;;
  fish)
    FISH_CONFIG="$HOME/.config/fish/config.fish"
    echo "alias itcm 'imotif-tools commit'" >> "$FISH_CONFIG"
    echo "Alias added to $FISH_CONFIG"
    ;;
  *)
    echo "Unknown shell: $SHELL_NAME. Please add alias manually:"
    echo "alias itcm='imotif-tools commit'"
    ;;
esac

echo "Please restart your terminal or run 'source' on your shell config to use 'itcm'"
