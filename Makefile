build:
	go build

run:
	env session_hash=needed_hash dbname=cetch_production base_path=$$GOPATH/src/github.com/JacksonGariety/cetch ./cetch

install:
	go get github.com/tools/godep
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
	env session_hash=dev_hash dbname=cetch_development base_path=$$GOPATH/src/github.com/JacksonGariety/cetch fresh

test:
	dropdb cetch_test --if-exists
	createdb cetch_test
	goose -env test up
	env session_hash=test_hash dbname=cetch_test base_path=$$GOPATH/src/github.com/JacksonGariety/cetch/ go test ./app/...
	dropdb cetch_test
