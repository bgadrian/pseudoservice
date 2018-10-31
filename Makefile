# Makefile
source := main.go

pre:
	mkdir -p ./build/
	env GO111MODULE=on go get -d ./

run: pre
	go run $(source)

build: pre
	go build -o ./build/server $(source)
	@echo "See ./build/server --help"

buildall: pre
	mkdir -p ./build/pseudoservice/windows
	mkdir -p ./build/pseudoservice/linux
	mkdir -p ./build/pseudoservice/macos
	GOOS=darwin GOARCH=amd64 go build -o ./build/pseudoservice/macos/pseudoservice $(source)
	GOOS=linux GOARCH=amd64 go build -o ./build/pseudoservice/linux/pseudoservice $(source)
	GOOS=windows GOARCH=amd64 go build -o  ./build/pseudoservice/windows/pseudoservice.exe $(source)
	cd ./build && tar -czf ./pseudoservice.tar.gz ./pseudoservice/
	@echo "publish to gihub: $ hub release create -a ./build/pseudoservice.tar.gz -m 'v0.X' v0.X"