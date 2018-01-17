package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Target struct {
	Name string
	Url  string
}

func (t *Target) Fire(f Feed) error {
	fmt.Println(t.Name, " Fire ", f.Title)
	form := url.Values{}
	form.Add("url", f.Link)
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(t.Url, "application/x-www-form-urlencoded", body)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	var resp map[string]interface{}

	if err := json.Unmarshal(body_byte, &resp); err != nil {
		return err
	}

	fmt.Printf("%s [%s] %s \n", t.Name, resp["status"], f.Title)
	return nil
}
