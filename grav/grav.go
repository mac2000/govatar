package grav

import (
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
)

//go:embed requests
var requests embed.FS

type Gravatar struct {
	username string
	password string
	api      string
}

func NewGravatarClient(username string, password string) Gravatar {
	return Gravatar{
		username: username,
		password: password,
		api:      buildApiUrl(username),
	}
}

func EmailHash(email string) string {
	b := md5.Sum([]byte(email))
	return hex.EncodeToString(b[:])
}

func buildApiUrl(username string) string {
	return fmt.Sprintf("https://secure.gravatar.com/xmlrpc?user=%s", EmailHash(username))
}

type Rating int

const (
	// 0:g, 1:pg, 2:r, 3:x
	RatingG = 0
	RatingP = 1
	RatingR = 2
	RatingX = 3
)
