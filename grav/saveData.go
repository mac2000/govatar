package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SaveData interface {
	SaveData() (int, error)
}

func (g Gravatar) SaveData(rating Rating, imageBase64 string) (string, error) {
	b, err := requests.ReadFile("requests/saveData.xml")
	if err != nil {
		return "", err
	}

	requestXmlString := fmt.Sprintf(string(b), imageBase64, rating, g.password)
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
	v := saveDataResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return "", err
	}

	if v.Error != "" {
		return "", errors.New(v.Error)
	}

	return v.UserImage, nil
}

type saveDataResponseXml struct {
	UserImage string `xml:"params>param>value>string"`
	Error     string `xml:"fault>value>struct>member>value>string"`
}
