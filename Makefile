SERVICENAME=image_packages
SERVICEURL=github.com/ManyakRus/$(SERVICENAME)

FILEMAIN=./internal/main.go
FILEAPP=./bin/image_packages

NEW_REPO=github.com/ManyakRus/image_packages


run:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	#	cd ./bin && \
	./bin/app_race
mod:
	clear
	go get -u ./...
	go mod tidy -compat=1.18
	go mod vendor
	go fmt ./...
build:
	clear
	go build -race -o $(FILEAPP) $(FILEMAIN)
	cd ./cmd && \
	./VersionToFile.py
	cp $(FILEAPP) $(GOPATH)/bin
lint:
	clear
	go fmt ./...
	golangci-lint run ./internal/v0/...
	golangci-lint run ./pkg/v0/...
	gocyclo -over 10 ./internal/v0
	gocyclo -over 10 ./pkg/v0
	gocritic check ./internal/v0/...
	gocritic check ./pkg/v0/...
	staticcheck ./internal/v0/...
	staticcheck ./pkg/v0/...
run.test:
	clear
	go fmt ./...
	go test -coverprofile cover.out ./internal/v0/app/...
	go tool cover -func=cover.out
newrepo:
	sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g' go.mod
	find -name *.go -not -path "*/vendor/*"|xargs sed -i 's+$(SERVICEURL)+$(NEW_REPO)+g'
graph:
	clear
	image_packages ./ docs/packages.graphml

conn:
	clear
	image_connections ./internal docs/connections.graphml $(SERVICENAME)
