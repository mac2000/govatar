package grav

import (
	"testing"
)

func TestSaveUrlSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/saveUrl.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	userImage, err := g.SaveUrl(RatingG, "http://placehold.it/200")
	if err != nil {
		t.Error(err)
	}
	if userImage != "222" {
		t.Errorf("expected 222 got %s", userImage)
	}
}

func TestSaveUrlError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	userImage, err := g.SaveUrl(RatingG, "http://placehold.it/200")
	if err == nil {
		t.Error("expected error, got nil")
	}
	if userImage != "" {
		t.Errorf("expected emptry string, got %s", userImage)
	}
}
