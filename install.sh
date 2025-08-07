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
