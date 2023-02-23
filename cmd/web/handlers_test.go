package main

import (
	"github.com/Faisal135711/snippetbox/internal/assert"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, string(body), "OK")
}
