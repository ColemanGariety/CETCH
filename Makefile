run:
	go build
	env session_hash=needed_hash dbname=cetch env=production base_path=$$GOPATH/src/github.com/JacksonGariety/cetch ./cetch

install:
	go get github.com/tools/godep
	go get bitbucket.org/liamstask/goose/cmd/goose
	godep restore

install_dev:
	go get github.com/pilu/fresh
	go get github.com/stretchr/testify

install_production_arch:
	sudo pacman -U mbox/mbox-20130226-1-x86_64.pkg.tar.xz

db_setup_production:
	createuser cetch -s -U postgres
	createdb cetch -U cetch

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
