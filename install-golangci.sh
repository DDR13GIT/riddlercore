#! /bin/sh

# Check if golangci-lint is already installed
if ! [ -x "$(command -v golangci-lint)" ]; then
    echo "Installing golangci-lint linter..."

    # Using the latest official v2 installation script
    # This automatically detects your OS and architecture
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v2.8.0

    # Move binary to local bin and cleanup
    sudo mv bin/golangci-lint /usr/local/bin/
    rm -rf bin

    echo "Successfully installed golangci-lint v2.8.0!"
else
    echo "golangci-lint is already installed: $(golangci-lint --version)"
fi
