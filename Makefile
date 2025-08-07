.PHONY: build run clean install

build:
	go build -o yt-fastmeta main.go

run:
	go run main.go

install:
	go install

clean:
	rm -f yt-fastmeta
