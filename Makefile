default: build

build: dist
	go run cli/main.go

clean:
	rm -r dist

dist:
	@mkdir -p dist

.PHONY: clean
