build:
	go build -ldflags="-X 'go-file-share/configs.Version=development'" -o ./file-share main.go

build/linux:
	GOARCH=amd64 GOOS=linux go build -ldflags="-X 'go-file-share/configs.Version=development'" -trimpath -o ./file-share main.go

build/windows:
	GOARCH=amd64 GOOS=windows go build -ldflags="-X 'go-file-share/configs.Version=development'" -trimpath -o ./file-share.exe main.go

build/darwin:
	GOARCH=arm64 GOOS=darwin go build -ldflags="-X 'go-file-share/configs.Version=development'" -trimpath -o ./file-share main.go

run: 
	go run main.go --dir=~ -R -p 8080

clean:
	rm -f file-share
	rm -f file-share.exe

tidy:
	go mod tidy

