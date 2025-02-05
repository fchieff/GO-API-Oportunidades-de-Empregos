
.PHONY: default run build test docs clean

#Variables
APP_NAME=GO-API-Oportunidades-de-Empregos.git

#Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go --docs

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./ ...

docs:
	swag init

clean:	
	@rm -rf $(APP_NAME)
	@rm -rf ./docs