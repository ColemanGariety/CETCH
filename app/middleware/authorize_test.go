package middleware

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"bytes"
)


func makeTestHandlerFail(t *testing.T) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user, ok := CurrentUser(r)
		assert.False(t, ok)
		assert.Empty(t, user.Username)
	}
	return http.HandlerFunc(fn)
}

func TestAuthorizeFail(t *testing.T) {
	ts := httptest.NewServer(Authorize(makeTestHandlerFail(t)))
	defer ts.Close()

	var u bytes.Buffer
	u.WriteString(string(ts.URL))
	u.WriteString("/")

	res, err := http.Get(u.String())
	assert.NoError(t, err)
	if res != nil {
		defer res.Body.Close()
	}
}
