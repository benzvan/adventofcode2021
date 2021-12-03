default: lint fmt vet test build

build:
	go build -o build/$(binname) ./day01/day01.1
	go build -o build/$(binname) ./day01/day01.2
	go build -o build/$(binname) ./day02/day02.1

clean:
	rm -rf build

fmt:
	go fmt ./...

lint:
	golint ./...

vet:
	go vet ./...

test:
	mkdir -p docs
	go test -coverprofile=coverage.out ./...
	go tool cover -func ./coverage.out > coverage.txt
	go tool cover -html=coverage.out -o docs/coverage.html
	rm coverage.out


.PHONY: all build clean install lint fmt test vet