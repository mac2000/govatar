package grav

import (
	"testing"
)

func TestDeleteUserImageSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/deleteUserimage.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	err = g.DeleteUserimage("5f3da2561611ab7c88eb919b3345d00c")
	if err != nil {
		t.Error("expected nil got ", err.Error())
	}
}

func TestDeleteUserImageError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	err = g.DeleteUserimage("5f3da2561611ab7c88eb919b3345d00c")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
