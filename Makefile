.PHONY: build clean deploy gomodgen
compile = env GOOS=linux  GOARCH=amd64  go build -v -ldflags '-s -w -v' -o

build: gomodgen
	export GO111MODULE=on
	(compile) bin/hello hello/main.go
	(compile) bin/world world/main.go
	(compile) bin/createUserHandler handler/CreateUserHandler.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
