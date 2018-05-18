CGO_ENABLED=0

all: windows linux darwin

linux:
	GOOS=linux GOARCH=amd64 go build -o extract-from-yaml-linux main.go

windows:
	GOOS=windows GOARCH=amd64 go build -o extract-from-yaml.exe main.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o extract-from-yaml-darwin main.go

clean:
	rm extract-from-yaml-linux
	rm extract-from-yaml.exe
	rm extract-from-yaml-darwin