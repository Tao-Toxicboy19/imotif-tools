#!/bin/bash

set -e

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

# Ask user for path
echo -n "Enter custom install path or press [Enter] to use default: "
read CUSTOM_PATH

# Fallback to default if input is empty or whitespace
if [[ -z "$CUSTOM_PATH" || "$CUSTOM_PATH" =~ ^[[:space:]]*$ ]]; then
    INSTALL_PATH="$DEFAULT_PATH"
else
    INSTALL_PATH="$CUSTOM_PATH"
fi

# Create directory if not exist
mkdir -p "$INSTALL_PATH"

# Download binary
curl -L "$BINARY_URL" -o "$INSTALL_PATH/itgc"
chmod +x "$INSTALL_PATH/itgc"

echo "itgc installed to: $INSTALL_PATH/itgc"

# Check if path in $PATH
if [[ ":$PATH:" != *":$INSTALL_PATH:"* ]]; then
    echo ""
    echo "$INSTALL_PATH is not in your PATH."
    echo "Add this line to your shell config file:"
    echo ""
    echo "  export PATH=\"\$PATH:$INSTALL_PATH\""
    echo ""
    echo "Then run: source ~/.bashrc | ~/.zshrc | ~/.config/fish/config.fish"
else
    echo "You can now run 'itgc'"
fi
