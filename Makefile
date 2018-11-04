# Makefile
source := ./cmd/pseudo-service-server/main.go

pre:
	mkdir -p ./build/
	env GO111MODULE=on go get -d ./
	env GO111MODULE=on go test -race ./...

run: pre
	env PORT=8080 go run $(source)

build: pre
	rm -f ./build/pseudoservice
	env GO111MODULE=on go build -o ./build/pseudoservice $(source)
	@echo "See ./build/pseudoservice --help"

buildall: pre
	mkdir -p ./build/pseudoservice/windows
	mkdir -p ./build/pseudoservice/linux
	mkdir -p ./build/pseudoservice/macos
	env GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o ./build/pseudoservice/macos/pseudoservice $(source)
	env GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o ./build/pseudoservice/linux/pseudoservice $(source)
	env GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o  ./build/pseudoservice/windows/pseudoservice.exe $(source)
	cd ./build && tar -czf ./pseudoservice.tar.gz ./pseudoservice/
	@echo "publish to gihub: $ hub release create -a ./build/pseudoservice.tar.gz -m 'v0.X' v0.X"