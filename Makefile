BUILDPATH=$(CURDIR)
API_NAME=api-bfcl-create-customer
PACKAGES := $(shell go list ./src/... | grep -vE 'mocks')

build:
	@echo "Creando Binario ..."
	@go mod download
	@go build -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test:
	@echo "Ejecutando tests..."
	@go test $(PACKAGES) --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

run: docker
	go run main.go

docker:
	docker-compose up -d