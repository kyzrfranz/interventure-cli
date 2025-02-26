# Variables
CMD_DIR := ./cmd
TMP_DIR := ./.tmp

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
	@rm -rf $(TMP_DIR)/*

bio-all:
	@echo "Building bio..."
	@go build -o ./bin/bio $(CMD_DIR)/bio.go
	@go run $(CMD_DIR)/bio.go

fetch-images:
	@echo "Downloading img..."
	@go build -o ./bin/bio $(CMD_DIR)/bio.go
	@go run $(CMD_DIR)/transform.go --type img

process-images: img-resize-150 img-resize-75 img-crop-square

img-resize-150:
	magick mogrify \
      -format webp \
      -strip \
      -resize 150x \
      -quality 80 \
      -set filename:out '%[basename]_medium' \
      -write '$(TMP_DIR)/%[filename:out].webp' \
      +delete \
      $(TMP_DIR)/*.jpg

img-resize-75:
	magick mogrify \
		  -format webp \
		  -strip \
		  -resize 75x \
		  -quality 80 \
		  -set filename:out '%[basename]_small' \
		  -write '$(TMP_DIR)/%[filename:out].webp' \
		  +delete \
		  $(TMP_DIR)/*.jpg

img-crop-square:
	magick mogrify \
          -format webp \
          -strip \
          -gravity center \
          -crop 150x150+0+0 \
          +repage \
          -quality 80 \
          -set filename:out '%[basename]_150x150' \
          -write '$(TMP_DIR)/%[filename:out].webp' \
          +delete \
          $(TMP_DIR)/*.jpg