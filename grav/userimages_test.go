package grav

import (
	"testing"
)

func TestUserImagesSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/userimages.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	userImages, err := g.UserImages()
	if err != nil {
		t.Error(err)
	}
	if len(userImages) != 2 {
		t.Errorf("expected 2 userImages, got %v", len(userImages))
	}
	assertUserImage(t, userImages[0], UserImage{
		Name:   "b005b5087e5761036271996cc2f76a89",
		Rating: RatingG,
		URL:    "http://en.gravatar.com/userimage/4299573/b005b5087e5761036271996cc2f76a89.jpg",
	})
	assertUserImage(t, userImages[1], UserImage{
		Name:   "07716eddaac23192e97bae292e95da33",
		Rating: RatingG,
		URL:    "http://en.gravatar.com/userimage/4299573/07716eddaac23192e97bae292e95da33.jpg",
	})
}

func TestUserImagesError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	userImages, err := g.UserImages()
	if err == nil {
		t.Error("expected error, got nil")
	}
	if len(userImages) != 0 {
		t.Errorf("expected 0 userImages, got %v", len(userImages))
	}
}

func assertUserImage(t *testing.T, actual UserImage, expected UserImage) {
	if actual.Name != expected.Name {
		t.Errorf("expected %s, got %s", expected.Name, actual.Name)
	}
	if actual.Rating != expected.Rating {
		t.Errorf("expected %v, got %v", expected.Rating, actual.Rating)
	}
	if actual.URL != expected.URL {
		t.Errorf("expected %s, got %s", expected.URL, actual.URL)
	}
}
