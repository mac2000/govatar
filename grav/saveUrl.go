package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SaveUrl interface {
	SaveUrl() (int, error)
}

func (g Gravatar) SaveUrl(rating Rating, imageURL string) (string, error) {
	b, err := requests.ReadFile("requests/saveUrl.xml")
	if err != nil {
		return "", err
	}

	requestXmlString := fmt.Sprintf(string(b), imageURL, rating, g.password)
	// fmt.Println(requestXmlString)
	r := bytes.NewReader([]byte(requestXmlString))
	resp, err := http.Post(g.api, "text/xml", r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(body))
	v := saveUrlResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return "", err
	}

	if v.Error != "" {
		return "", errors.New(v.Error)
	}

	return v.UserImage, nil
}

type saveUrlResponseXml struct {
	UserImage string `xml:"params>param>value>string"`
	Error     string `xml:"fault>value>struct>member>value>string"`
}
