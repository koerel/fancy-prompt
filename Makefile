steps := clean macos linux

.PHONY: $(steps)

default: $(steps)

clean:
	rm builds/*

macos:
	env GOOS=darwin GOARCH=amd64 go build -o builds/fancy-prompt-macos *.go

linux:
	env GOOS=linux GOARCH=amd64 go build -o builds/fancy-prompt-linux *.go