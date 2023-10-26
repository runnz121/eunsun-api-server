.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function1 ./functions/function1/main.go
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function2 ./functions/function2/main.go
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function3 ./functions/function3/main.go
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function4 ./functions/function4/main.go

build-function1:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function1 ./functions/function1/main.go

build-function2:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function2 ./functions/function2/main.go

build-function3:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function3 ./functions/function3/main.go

build-function4:
	env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/function4 ./functions/function4/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

local: build
	sls offline -s stage

deploy: clean build
	sls deploy -s $(STAGE)

deploy-function1: clean build-function1
	sls deploy -s $(STAGE) -f function1

deploy-function2: clean build-function2
	sls deploy -s $(STAGE) -f function2

deploy-function3: clean build-function3
	sls deploy -s $(STAGE) -f function3

deploy-function4: clean build-function4
	sls deploy -s $(STAGE) -f function4

