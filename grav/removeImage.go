package grav

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RemoveImage interface {
	RemoveImage() (int, error)
}

func (g Gravatar) RemoveImage(email string) error {
	b, err := requests.ReadFile("requests/removeImage.xml")
	if err != nil {
		return err
	}

	requestXmlString := fmt.Sprintf(string(b), email, g.password)
	// fmt.Println(requestXmlString)
	r := bytes.NewReader([]byte(requestXmlString))
	resp, err := http.Post(g.api, "text/xml", r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))
	v := removeImageResponseXml{}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return err
	}

	if v.Error != "" {
		return errors.New(v.Error)
	}

	if v.Success != 1 {
		return fmt.Errorf("unexpected result %v", v.Success)
	}

	return nil
}

type removeImageResponseXml struct {
	Success int    `xml:"params>param>value>struct>member>value>boolean"`
	Error   string `xml:"fault>value>struct>member>value>string"`
}
