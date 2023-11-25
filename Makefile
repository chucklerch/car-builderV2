# Make is verbose, go silent.
MAKEFLAGS += --silent
#SHELL=/bin/bash

# Defaults
SOURCES=$(wildcard *.go)            	   	# Source files
EXECUTABLE=car-builder                      # Program name

.PHONY: all
all: $(EXECUTABLE)  						# Build executable

$(EXECUTABLE): $(SOURCES) ## Create executable
	@echo "\033[33;1mBuilding\033[0m"
	go build .

.PHONY: clean
clean: ## Clean up
	@echo "\033[33;1mCleaning\033[0m"
	rm -f $(EXECUTABLE)

.PHONY: help
help: ## Display help
	@echo "Make targets"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo

