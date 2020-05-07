all: window linux

window: 
	GOOS=windows GOARCH=386 go build

linux: clean_linux
	@echo "Building file for linux..."
	GOOS=linux GOARCH=amd64 go build
	@chmod +x xi_video_downloader

clean_linux:
	rm -rf ./xi_video_downloader

.PHONY: window linux