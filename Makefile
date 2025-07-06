# Makefile for task-cli

BINARY_NAME = task-cli
INSTALL_DIR = $(HOME)/.local/bin

.PHONY: build install clean

build:
	go mod tidy
	go build -o $(BINARY_NAME)

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)/
	chmod +x $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME) to $(INSTALL_DIR)"
	@echo "Make sure $(INSTALL_DIR) is in your PATH."

clean:
	rm -f $(BINARY_NAME)
