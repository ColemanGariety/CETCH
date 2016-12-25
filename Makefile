build:
	go build

run:
	./cetch

install:
	godep update

install_dev:
	go get github.com/pilu/fresh
	go get bitbucket.org/liamstask/goose/cmd/goose

schema:
	goose up

watch:
	fresh
