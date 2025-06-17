.PHONY: build
build:
	go build -o bin/imgbase ./cmd/imgbase

.PHONY: clean
clean:
	rm -rf bin/