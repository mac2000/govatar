package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Test interface {
	Test() (int, error)
}

func (g Gravatar) Test() (int, error) {
	b, err := requests.ReadFile("requests/test.xml")
	if err != nil {
		return 0, err
	}

	requestXmlString := fmt.Sprintf(string(b), g.password)
	r := bytes.NewReader([]byte(requestXmlString))
	resp, err := http.Post(g.api, "text/xml", r)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	v := testResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return 0, err
	}

	if v.Error != "" {
		return 0, errors.New(v.Error)
	}

	return v.ID, nil
}

// func Test(username string, password string) (int, error) {
// 	api := "https://secure.gravatar.com/xmlrpc?user=%s"

// 	usernameBytes := md5.Sum([]byte(username))
// 	hash := hex.EncodeToString(usernameBytes[:])

// 	url := fmt.Sprintf(api, hash)

// 	b, err := requests.ReadFile("requests/test.xml")
// 	if err != nil {
// 		return 0, err
// 	}

// 	requestXmlString := fmt.Sprintf(string(b), password)

// 	r := bytes.NewReader([]byte(requestXmlString))
// 	resp, err := http.Post(url, "text/xml", r)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return 0, err
// 	}
// 	fmt.Println(string(body))
// 	v := testResponseXml{}

// 	err = xml.Unmarshal(body, &v)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if v.Error != "" {
// 		return 0, errors.New(v.Error)
// 	}

// 	return v.ID, nil
// }

type testResponseXml struct {
	ID    int    `xml:"params>param>value>struct>member>value>int"`
	Error string `xml:"fault>value>struct>member>value>string"`
}
