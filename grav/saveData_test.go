package grav

import (
	"testing"
)

func TestSaveDataSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/saveData.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	userImage, err := g.SaveData(RatingG, "base64ofimage==")
	if err != nil {
		t.Error(err)
	}
	if userImage != "111" {
		t.Errorf("expected 111 got %s", userImage)
	}
}

func TestSaveDataError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	userImage, err := g.SaveData(RatingG, "base64ofimage==")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if userImage != "" {
		t.Errorf("expected emptry string, got %s", userImage)
	}
}
