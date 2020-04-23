all: window linux

window: 
	GOOS=windows GOARCH=386 go build

linux: 
	GOOS=linux GOARCH=amd64 go build


.PHONY: window linux