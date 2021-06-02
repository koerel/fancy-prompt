steps := clean macos linux

.PHONY: $(steps)

default: $(steps)

clean:
	rm builds/* || true

macos:
	env GOOS=darwin GOARCH=amd64 go build -o builds/fancy-prompt-macos *.go

linux:
	env GOOS=linux GOARCH=amd64 go build -o builds/fancy-prompt-linux *.go

install:
	sudo cp builds/fancy-prompt-linux /usr/local/bin/fancy-prompt