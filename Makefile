default: build

build: dist
	go run .

clean:
	rm -r dist

dist:
	@mkdir -p dist

.PHONY: clean
