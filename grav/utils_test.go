package grav

import (
	"net/http"
	"net/http/httptest"
	"os"
)

func mockGravatar(output string) (*httptest.Server, Gravatar, error) {
	g := NewGravatarClient("test@gmail.com", "123")

	r, err := os.ReadFile(output)
	if err != nil {
		return nil, g, err
	}

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write(r)
	}))

	g.api = ts.URL

	return ts, g, nil
}
