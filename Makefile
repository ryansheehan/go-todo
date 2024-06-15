.PHONY: all run build_dev build_template

all: build_dev run

build_template:
	@templ generate

build_dev:
	@templ generate
	go build -o ./tmp/main ./cmd/...

run:
	@./tmp/main
