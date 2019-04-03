build:
	go build

test:
	go test ./ -v

fmt:
	go fmt ./

coverage:
	go test ./ -coverprofile=cover.out
	go tool cover -html=cover.out

clean:
	rm cover.out

lint:
	go vet ./

deps:
	go mod download