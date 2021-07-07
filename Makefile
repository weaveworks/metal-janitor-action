TOOLS_DIR := tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin

PATH := $(abspath $(TOOLS_BIN_DIR)):$(PATH)
export PATH

$(TOOLS_BIN_DIR):
	mkdir -p $@

GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint
GINKGO := $(TOOLS_BIN_DIR)/ginkgo

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run -v

$(GOLANGCI_LINT): go.mod # Get and build golangci-lint
	cd $(TOOLS_DIR); go build -tags=tools -o $(subst tools/,,$@) github.com/golangci/golangci-lint/cmd/golangci-lint