build:
	go build

run:
	./cetch

install:
	godep restore

install_dev:
	go get github.com/pilu/fresh
	go get bitbucket.org/liamstask/goose/cmd/goose
	go get github.com/stretchr/testify

db_status:
	goose status

db_migrate:
	goose -env development up

watch:
	env dbname=cetch_development basePath=$$GOPATH/src/github.com/JacksonGariety/cetch fresh

test:
	dropdb cetch_test --if-exists
	createdb cetch_test
	goose -env test up
	env dbname=cetch_test basePath=$$GOPATH/src/github.com/JacksonGariety/cetch/ go test ./app/...
	dropdb cetch_test
