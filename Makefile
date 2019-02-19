steps := clean macos linux deploy-macos deploy-linux

.PHONY: $(steps)

default: $(steps)

clean:
	rm builds/*

macos:
	env GOOS=darwin GOARCH=amd64 go build -o builds/fancy-prompt-macos *.go

linux:
	env GOOS=linux GOARCH=amd64 go build -o builds/fancy-prompt-linux *.go

deploy-macos:
	aws s3 cp builds/fancy-prompt-macos s3://bgr-assets/prompt/builds/fancy-prompt-macos --acl public-read

deploy-linux:
	aws s3 cp builds/fancy-prompt-macos s3://bgr-assets/prompt/builds/fancy-prompt-macos --acl public-read