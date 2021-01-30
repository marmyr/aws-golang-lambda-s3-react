del /Q bin
set GOOS=linux
go build -ldflags="-s -w" -o bin/aws endpoints/aws/aws.go