package grav

import (
	"testing"
)

func TestUseUserImageSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/useUserimage.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	err = g.UseUserImage("5f3da2561611ab7c88eb919b3345d00c", "test@test.com")
	if err != nil {
		t.Error("expected nil got ", err.Error())
	}
}

func TestUseUserImageError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	err = g.UseUserImage("5f3da2561611ab7c88eb919b3345d00c", "test@test.com")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
