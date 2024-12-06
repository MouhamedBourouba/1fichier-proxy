GO = go
PKG = .
BIN_DIR = bin
BINARY_NAME = 1fichier

# Default target
all: run

# Build the project
build:
	@echo "Building project..."
	$(GO) build -o $(BIN_DIR)/$(BINARY_NAME) $(PKG)

# Run the application
run: build
	@echo "Running application..."
	$(BIN_DIR)/$(BINARY_NAME)
