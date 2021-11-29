package grav

import (
	"testing"
)

func TestAddressesSuccess(t *testing.T) {
	ts, g, err := mockGravatar("responses/addresses.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	addresses, err := g.Addresses()
	if err != nil {
		t.Error(err)
	}
	if len(addresses) != 2 {
		t.Errorf("expected 2 addresses, got %v", len(addresses))
	}
	assertAddress(t, addresses[0], Address{
		Email:        "marchenko.alexandr@gmail.com",
		Rating:       Rating(0),
		UserImage:    "1",
		UserImageURL: "http://en.gravatar.com/userimage/a/b.jpg",
	})
	assertAddress(t, addresses[1], Address{
		Email:        "alexandrm@rabota.ua",
		Rating:       Rating(0),
		UserImage:    "2",
		UserImageURL: "http://en.gravatar.com/userimage/x/y.jpg",
	})
}

func TestAddressesError(t *testing.T) {
	ts, g, err := mockGravatar("responses/fault.xml")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()
	addresses, err := g.Addresses()
	if err == nil {
		t.Error("expected error, got nil")
	}
	if len(addresses) != 0 {
		t.Errorf("expected 0 addresses, got %v", len(addresses))
	}
}

func assertAddress(t *testing.T, actual Address, expected Address) {
	if actual.Email != expected.Email {
		t.Errorf("expected %s, got %s", expected.Email, actual.Email)
	}
	if actual.Rating != expected.Rating {
		t.Errorf("expected %v, got %v", expected.Rating, actual.Rating)
	}
	if actual.UserImage != expected.UserImage {
		t.Errorf("expected %s, got %s", expected.UserImage, actual.UserImage)
	}
	if actual.UserImageURL != expected.UserImageURL {
		t.Errorf("expected %s, got %s", expected.UserImageURL, actual.UserImageURL)
	}
}
