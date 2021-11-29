package grav

import (
	"testing"
)

func TestRemoveImageSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/removeImage.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	err = g.RemoveImage("5f3da2561611ab7c88eb919b3345d00c")
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveImageError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	err = g.RemoveImage("5f3da2561611ab7c88eb919b3345d00c")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
