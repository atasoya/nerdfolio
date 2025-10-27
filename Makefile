.PHONY: build run clean

build:
	go build -o bin/nerdfolio ./cmd/nerdfolio

build-test:
	go build -o bin/nerdfolio ./cmd/nerdfolio
	sudo ln -sf $(PWD)/bin/nerdfolio /usr/local/bin/nerdfolio




run: build
	./bin/nerdfolio

clean:
	rm -rf bin/
