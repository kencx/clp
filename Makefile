.PHONY: build

build:
	go build -ldflags '-s -w' -o clp .

clean:
	go clean
	rm -rf clp
