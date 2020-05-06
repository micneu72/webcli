export GOOS=darwin
export GOARCH=amd64
go build -o build/macOS/
export GOOS=windows
export GOARCH=amd64
go build -o build/win/
export GOOS=linux
export GPARCH=amd64
go build -o build/linux/
