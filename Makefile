

build_win:
	GOOS=windows GOARCH=amd64 go build -o bin/proxy.exe
