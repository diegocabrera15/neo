compile = env GOOS=linux  GOARCH=amd64  go build -v -ldflags '-s -w -v' -o
.PHONY: build clean deploy gomodgen

build: gomodgen
	go mod download github.com/aws/aws-lambda-go
	go get github.com/go/neo/service
	go get github.com/go/neo/connection
	go get github.com/go/neo/service
	export GO111MODULE=on
	$(compile) bin/createUserHandler handler/CreateUserHandler.go
	$(compile) bin/createUnitHandler handler/CreateUnitHandler.go
	$(compile) bin/readUnitHandler handler/ReadUnitHandler.go

clean:
	clear
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --aws-profile fer

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
