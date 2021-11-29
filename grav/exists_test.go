package grav

import (
	"testing"
)

func TestExistsSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/exists.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	exists, err := g.Exists("5f3da2561611ab7c88eb919b3345d00c")
	if err != nil {
		t.Error(err)
	}
	if exists != true {
		t.Errorf("expected true got false")
	}
}

func TestExistsError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	exists, err := g.Exists("5f3da2561611ab7c88eb919b3345d00c")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if exists == true {
		t.Errorf("expected false, got true")
	}
}
