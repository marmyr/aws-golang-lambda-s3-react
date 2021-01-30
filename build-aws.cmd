del /Q bin
set GOOS=linux
go build -ldflags="-s -w" -o bin/hello endpoints/aws/hello.go