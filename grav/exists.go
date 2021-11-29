package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Exists interface {
	Exists() (int, error)
}

func (g Gravatar) Exists(emailHash string) (bool, error) {
	b, err := requests.ReadFile("requests/exists.xml")
	if err != nil {
		return false, err
	}

	requestXmlString := fmt.Sprintf(string(b), emailHash, g.password)

	r := bytes.NewReader([]byte(requestXmlString))
	resp, err := http.Post(g.api, "text/xml", r)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	// fmt.Println(string(body))
	v := existsResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return false, err
	}

	if v.Error != "" {
		return false, errors.New(v.Error)
	}

	if v.Exists == 1 {
		return true, nil
	} else {
		return false, nil
	}
}

type existsResponseXml struct {
	Exists int    `xml:"params>param>value>struct>member>value>int"`
	Error  string `xml:"fault>value>struct>member>value>string"`
}
