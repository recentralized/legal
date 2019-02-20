default: gen

gen: build dist
	bin/legal -w

serve: build
	bin/legal -s

clean:
	rm -r dist
	rm -r bin

build:
	go build -o bin/legal cli/main.go

dist:
	@mkdir -p dist

.PHONY: build clean gen serve
