AUTHOR=rjfarmer
NAME=zkill-mirror
VERSION=0.1

.PHONY: build tag_latest clean

default: build tag_latest
    
# GOOS=linux GOARCH=amd64
build:
	CGO_ENABLED=0 go build -o zkill-mirror
	docker build -t $(AUTHOR)/$(NAME):$(VERSION) .

tag_latest:
	docker tag $(AUTHOR)/$(NAME):$(VERSION) $(AUTHOR)/$(NAME):latest

clean:
	go clean

