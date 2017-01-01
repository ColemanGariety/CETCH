install:
	go get github.com/tools/godep
	go get bitbucket.org/liamstask/goose/cmd/goose
	godep restore
	git clone https://github.com/projectatomic/bubblewrap.git
	cd bubblewrap && exec ./autogen.sh
	sudo $(MAKE) install -C bubblewrap
	rm -rf bubblewrap

install_dev:
	go get github.com/tools/godep
	go get bitbucket.org/liamstask/goose/cmd/goose
	godep restore
	go get github.com/pilu/fresh
	go get github.com/stretchr/testify

watch:
	env session_hash=dev_hash dbname=cetch_development base_path=$$GOPATH/src/github.com/JacksonGariety/cetch fresh

test:
	dropdb cetch_test --if-exists
	createdb cetch_test
	goose -env test up
	env session_hash=test_hash dbname=cetch_test base_path=$$GOPATH/src/github.com/JacksonGariety/cetch/ go test ./app/...
	dropdb cetch_test

run:
	go build
	env session_hash=needed_hash dbname=cetch env=production base_path=$$GOPATH/src/github.com/JacksonGariety/cetch ./cetch
