package firefly

import (
	"net/http"
	"log"
	"net/url"
	"encoding/json"
	"errors"
	"os"
	"io"
	"bytes"
)

type Client struct {
	log          *log.Logger
	http         *http.Client
	baseUrl      string
	secretApiKey string
	debugLog     bool
}

func NewClient(secretApiKey string) *Client {
	client := &Client{}
	client.http = http.DefaultClient
	client.log = log.New(os.Stderr, "", log.LstdFlags)
	client.secretApiKey = secretApiKey
	client.baseUrl = "https://api.fireflyiot.com/api/v1/"
	return client
}

func (client *Client) SetDebugLog(val bool) {
	client.debugLog = val
}

func (client Client) Url() (u *url.URL) {
	parsedUrl, err := url.Parse(client.baseUrl)
	u = parsedUrl

	if err != nil {
		client.log.Fatal("Failed to parse URL: " + client.baseUrl)
	}

	q := u.Query()
	q.Set("auth", client.secretApiKey)
	u.RawQuery = q.Encode()
	return u
}

func (client Client) getAndDecode(reqUrl *url.URL, target interface{}) (err error) {
	resp, err := client.get(reqUrl.String())

	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status + " - GET " + reqUrl.String())
	}

	return decodeJsonResponse(resp, &target)
}

func (client Client) postAndDecode(reqUrl *url.URL, request interface{}, target interface{}) (err error) {
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := client.post(reqUrl.String(), "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status + " - POST " + reqUrl.String())
	}

	return decodeJsonResponse(resp, &target)
}

func (client Client) putAndDecode(reqUrl *url.URL, request interface{}, target interface{}) (err error) {
	jsonBytes, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := client.put(reqUrl.String(), "application/json", bytes.NewReader(jsonBytes))
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(resp.Status + " - PUT " + reqUrl.String())
	}

	return decodeJsonResponse(resp, &target)
}

func (client Client) deleteNoContent(reqUrl *url.URL) (err error) {
	resp, err := client.delete(reqUrl.String())

	if resp.StatusCode != 204 {
		return errors.New(resp.Status + " - DELETE " + reqUrl.String())
	}
	return
}

func decodeJsonResponse(resp *http.Response, target interface{}) (err error) {
	decoder := json.NewDecoder(resp.Body)
	if !decoder.More() {
		return errors.New("Received empty body")
	}

	err = decoder.Decode(&target)
	return
}

//////////////////////
// Basic http methods
//////////////////////

// like http.Post but for put, it's missing in http package
func (client *Client) put(url string, bodyType string, body io.Reader) (res *http.Response, err error) {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	res, err = client.http.Do(req)
	if client.debugLog {
		client.log.Println("PUT", res.Status, url)
	}

	return
}

func (client *Client) post(url string, bodyType string, body io.Reader) (res *http.Response, err error) {
	res, err = client.http.Post(url, bodyType, body)
	if client.debugLog {
		client.log.Println("POST", res.Status, url)
	}
	return
}

func (client *Client) get(url string) (res *http.Response, err error) {
	res, err = client.http.Get(url)
	if client.debugLog {
		client.log.Println("GET", res.Status, url)
	}
	return
}

func (client *Client) delete(url string) (res *http.Response, err error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	res, err = client.http.Do(req)
	if client.debugLog {
		client.log.Println("DELETE", res.Status, url)
	}
	return
}
