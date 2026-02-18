build:
	go build -o ./file-share main.go

build/linux:
	GOARCH=amd64 GOOS=linux go build -trimpath -o ./file-share main.go

build/windows:
	GOARCH=amd64 GOOS=windows go build -trimpath -o ./file-share.exe main.go

build/darwin:
	GOARCH=arm64 GOOS=darwin go build -trimpath -o ./file-share main.go

run: 
	go run main.go --dir=~ -R -p 8080

clean:
	rm -f file-share
	rm -f file-share.exe
