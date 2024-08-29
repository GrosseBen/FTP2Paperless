all: build-x86

run:
	docker run --rm -it \
    -e FTP_SERVER=192.168.112.155 \
    -e FTP_USER=consume \
    -e FTP_PASS=consume \
    -e REMOTE_DIR=/ \
    -e API_URL=http://example.com/api/documents/post_document/ \
    -e API_TOKEN=your_api_token_here \
    ludal/ftp2paperless

build:
	go build -o bin/ftp2paperless ./cmd/ftp2paperless

build-x86:
	docker build -t ludal/ftp2paperless .
