.PHONY: build clean run

clean:
	rm -rf ./bin

test:
	go test ./pkg/...

integration:
	go test --tags=integration ./pkg/...

build: clean integration
	for CMD in `ls cmd`; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$CMD ./cmd/$$CMD/...; \
	done
