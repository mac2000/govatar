package grav

import (
	"testing"
)

// var g = NewGravatarClient("test@gmail.com", "123")

func TestSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/test.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	id, err := g.Test()
	if err != nil {
		t.Error(err)
	}
	if id != 2 {
		t.Errorf("expected 2 got %v", id)
	}
}

func TestError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	id, err := g.Test()
	if err == nil {
		t.Error("expected error, got nil")
	}
	if id != 0 {
		t.Errorf("expected 0, got %v", id)
	}
}

// func TestSuccess(t *testing.T) {
// 	r, err := os.ReadFile("responses/test.xml")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		res.WriteHeader(200)
// 		res.Write(r)
// 	}))
// 	defer ts.Close()

// 	g.api = ts.URL
// 	id, err := g.Test()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if id != 2 {
// 		t.Errorf("expected 2 got %v", id)
// 	}
// }

// func TestError(t *testing.T) {
// 	r, err := os.ReadFile("responses/fault.xml")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
// 		res.WriteHeader(200)
// 		res.Write(r)
// 	}))
// 	defer ts.Close()

// 	g.api = ts.URL
// 	id, err := g.Test()
// 	if err == nil {
// 		t.Error("expected error, got nil")
// 	}
// 	if id != 0 {
// 		t.Errorf("expected 0, got %v", id)
// 	}
// }
