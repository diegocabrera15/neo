compile = env GOOS=linux  GOARCH=amd64  go build -v -ldflags '-s -w -v' -o
.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	$(compile) bin/createUserHandler handler/CreateUserHandler.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
