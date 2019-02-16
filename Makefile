default: build

build: dist
	go run main.go

clean:
	rm -r dist

dist:
	@mkdir -p dist

.PHONY: clean
