# Variables
CMD_DIR := ./cmd

# Targets
CMDS := bio transform

# Build and Run rules
.PHONY: all build clean run dev-run

all: build

build: $(CMDS)

$(CMDS):
	@echo "Building $@..."
	@go build -o ./bin/$@ $(CMD_DIR)/$@.go

run-%: build
	@./bin/$*

dev-run-%:
	@go run $(CMD_DIR)/$*.go

clean:
	@echo "Cleaning up..."
	@rm -rf ./bin
