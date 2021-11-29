package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Addresses interface {
	Addresses() ([]Address, error)
}

func (g Gravatar) Addresses() ([]Address, error) {
	b, err := requests.ReadFile("requests/addresses.xml")
	if err != nil {
		return nil, err
	}
	requestXmlString := fmt.Sprintf(string(b), g.password)

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

	v := addressesResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	if v.Error != "" {
		return nil, errors.New(v.Error)
	}

	addresses := []Address{}
	for _, member := range v.Members {
		addresses = append(addresses, Address{
			Email:        member.Name,
			Rating:       Rating(member.Values[0].Value.Int),
			UserImage:    member.Values[1].Value.String,
			UserImageURL: member.Values[2].Value.String,
		})
	}
	return addresses, nil
}

type addressesResponseXml struct {
	Members []struct {
		Name   string `xml:"name"`
		Values []struct {
			Name  string `xml:"name"`
			Value struct {
				Int    int    `xml:"int"`
				String string `xml:"string"`
			} `xml:"value"`
		} `xml:"value>struct>member"`
	} `xml:"params>param>value>struct>member"`

	Error string `xml:"fault>value>struct>member>value>string"`
}

type Address struct {
	Email        string
	Rating       Rating
	UserImage    string
	UserImageURL string
}
