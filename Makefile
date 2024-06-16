.PHONY: all run build build_js build_dev build_template

all: build_template build_js build_dev

build_template:
	@templ generate

build_dev:	
	go build -o ./tmp/main ./cmd/...

build_js:
	npm run build

run:
	@./tmp/main
