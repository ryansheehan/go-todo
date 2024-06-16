.PHONY: all run build build_dev build_template

all: build_template build_dev

build_template:
	@templ generate

build_dev:	
	go build -o ./tmp/main ./cmd/...

run:
	@./tmp/main
