# Makefile for Go practice workspace
# - vnv: create a Python virtualenv at .venv (useful for helper scripts/tools)
# - install-go: attempts to install Go via apt (Linux) if missing
# - setup: install-go + mod tidy
# - mod-tidy, fmt, vet, test, run-playground, new-problem, clean, help

.PHONY: venv install-go setup mod-tidy fmt vet test run-playground new-problem clean help

VENV_DIR := .venv

# Default go version hint (used only in messages). For specific Go install use manual steps.
GO_VERSION ?= 1.20

help:
	@echo "Available targets:"
	@echo "  venv            - create a Python virtualenv in ${VENV_DIR} (python3 required)"
	@echo "  install-go      - install Go (attempts apt install golang-go on Linux)"
	@echo "  setup           - install-go + go mod tidy"
	@echo "  mod-tidy        - run 'go mod tidy'"
	@echo "  fmt             - run 'go fmt ./...'"
	@echo "  vet             - run 'go vet ./...'"
	@echo "  test            - run 'go test ./...'"
	@echo "  run-playground  - run the playground (go run ./playground)"
	@echo "  new-problem     - create a new problem template (usage: make new-problem name=two_sum)"
	@echo "  clean           - remove ${VENV_DIR} and temporary files"

venv:
	@echo "Creating Python virtualenv at ${VENV_DIR}..."
	@if command -v python3 >/dev/null 2>&1; then \
		python3 -m venv ${VENV_DIR} && echo "Created ${VENV_DIR}"; \
	else \
		echo "python3 not found; please install Python 3 to use this target"; exit 1; \
	fi

install-go:
	@echo "Checking for 'go' in PATH..."
	@if command -v go >/dev/null 2>&1; then \
		echo "go already installed: $$(go version)"; \
	else \
		if command -v apt >/dev/null 2>&1; then \
			echo "Attempting to install 'golang-go' via apt (requires sudo)..."; \
			sudo apt update && sudo apt install -y golang-go; \
			if command -v go >/dev/null 2>&1; then \
				echo "Installed go: $$(go version)"; \
			else \
				echo "apt install finished but 'go' still not found. Please install Go from https://go.dev/dl/ manually."; exit 1; \
			fi; \
		else \
			echo "Automatic install not implemented for this OS. Visit https://go.dev/dl/ to install Go ${GO_VERSION} or newer."; exit 1; \
		fi; \
	fi

setup: install-go mod-tidy
	@echo "Setup complete (Go + mod tidy)."

mod-tidy:
	@echo "Running 'go mod tidy'..."
	@go mod tidy

fmt:
	@echo "Running 'go fmt ./...'"
	@go fmt ./...

vet:
	@echo "Running 'go vet ./...'"
	@go vet ./...

test:
	@echo "Running 'go test ./...'"
	@go test ./...

run-playground:
	@echo "Running playground..."
	@go run ./playground



new-problem:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make new-problem name=problem_name"; exit 1; \
	fi
	mkdir -p problems
	./scripts/new_problem.sh "$(name)"
	@echo "Done."
remove-problem:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make remove-problem name=problem_name"; exit 1; \
	fi
	rm -f problems/$(name).go
	@echo "Removed problems/$(name).go"
	@echo "Done."

clean:
	@echo "Cleaning..."
	@rm -rf ${VENV_DIR}
	@echo "Removed ${VENV_DIR}" || true
	@echo "Done."
