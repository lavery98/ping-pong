IMAGE_NAME := lavery98/ping-pong

.PHONY: build
build:
	CGO_ENABLED=0 go build -a --trimpath --installsuffix cgo --ldflags '-s' -o ping-pong

.PHONY: image
image:
	docker build -t ${IMAGE_NAME} .
