.PHONY: build run clean

build:
	go build -o bin/nerdfolio ./cmd/nerdfolio

run: build
	./bin/nerdfolio

clean:
	rm -rf bin/
