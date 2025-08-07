#!/bin/bash

set -e

# Detect OS
OS=$(uname)
BINARY_URL=""

if [[ "$OS" == "Darwin" ]]; then
    BINARY_URL="https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/itgc-macos"
elif [[ "$OS" == "Linux" ]]; then
    BINARY_URL="https://github.com/Tao-Toxicboy19/imotif-tools/releases/latest/download/itgc-linux"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Default install path
DEFAULT_PATH="$HOME/.local/bin"
echo "Default install path is: $DEFAULT_PATH"

# Ask for custom path
read -p "Enter custom install path or press [Enter] to use default: " CUSTOM_PATH
INSTALL_PATH="${CUSTOM_PATH:-$DEFAULT_PATH}"

# Create dir if not exist
mkdir -p "$INSTALL_PATH"

# Download binary
curl -L "$BINARY_URL" -o "$INSTALL_PATH/itgc"
chmod +x "$INSTALL_PATH/itgc"

echo "itgc installed to: $INSTALL_PATH/itgc"

# Check if install path is in PATH
if [[ ":$PATH:" != *":$INSTALL_PATH:"* ]]; then
    echo ""
    echo "$INSTALL_PATH is not in your PATH."
    echo "Add the following line to your shell config file:"
    echo ""
    echo "  export PATH=\"\$PATH:$INSTALL_PATH\""
    echo ""
    echo "For example:"
    echo "  ~/.bashrc, ~/.zshrc, ~/.config/fish/config.fish (fish: set -Ux PATH \$PATH $INSTALL_PATH)"
    echo ""
    echo "After that, restart your terminal or run: source <your-shell-file>"
else
    echo "You can now run 'itgc' from your terminal!"
fi
