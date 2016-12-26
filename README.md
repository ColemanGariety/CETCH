# cetch.com

Install the deps (`godep restore`):

    $ make install

Godeps doesn't handle dev deps so we do this to install dev tools (fresh, goose, and testify):

    $ make install_dev

You'll want to create the postgres role and development database:

    $ sudo -i -u postgres
    $ createuser -s cetch
    $ createdb cetch_development

Then run the migrations via goose (with the `-env development` flag):

    $ make db_migrate

Watch for changes while working (sets `dbname` and `base_path` flags):

    $ make watch

And run the test suite (test db is automatically created before tests run and dropped afterwards):

    $ make test

Running the application in production without make would look like this:

    $ go build
    $ env dbname=cetch_production basePath=$GOPATH/src/github.com/JacksonGariety/cetch ./cetch
