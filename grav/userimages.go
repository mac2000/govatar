package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type UserImages interface {
	UserImages() ([]Address, error)
}

func (g Gravatar) UserImages() ([]UserImage, error) {
	b, err := requests.ReadFile("requests/userimages.xml")
	if err != nil {
		return nil, err
	}
	// requestXmlString := fmt.Sprintf(string(b), rating, password)
	requestXmlString := fmt.Sprintf(string(b), g.password)
	// fmt.Println(requestXmlString)

	r := bytes.NewReader([]byte(requestXmlString))
	resp, err := http.Post(g.api, "text/xml", r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// log.Println(string(body))

	v := userImagesResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	if v.Error != "" {
		return nil, errors.New(v.Error)
	}

	images := []UserImage{}
	for _, member := range v.Members {
		rating, _ := strconv.Atoi(member.Value[0].String)
		images = append(images, UserImage{
			Name:   member.Name,
			Rating: Rating(rating),
			URL:    member.Value[1].String,
		})
	}
	// b, _ = json.MarshalIndent(images, "", "  ")
	// fmt.Println(string(b))
	return images, nil
}

type userImagesResponseXml struct {
	Members []struct {
		Name  string `xml:"name"`
		Value []struct {
			String string `xml:"string"`
		} `xml:"value>array>data>value"`
	} `xml:"params>param>value>struct>member"`

	Error string `xml:"fault>value>struct>member>value>string"`
}

type UserImage struct {
	Name   string
	Rating Rating
	URL    string
}
