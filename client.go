package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
)

// Endpoint URL
const Endpoint = "http://api.tushare.pro"

// TuShare instance
type TuShare struct {
	token  string
	client *http.Client
}

// New TuShare default client
func New(token string) *TuShare {
	return NewWithClient(token, http.DefaultClient)
}

// NewWithClient TuShare client with arguments
func NewWithClient(token string, httpClient *http.Client) *TuShare {
	return &TuShare{
		token:  token,
		client: httpClient,
	}
}

func (api *TuShare) request(method, path string, body interface{}) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (api *TuShare) doRequest(req *http.Request) (*APIResponse, error) {
	// Set http content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := api.client.Do(req)
	//Handle network error
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("oops! Network error")
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read request
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check mime type of response
	mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	if mimeType != "application/json" {
		return nil, fmt.Errorf("Could not execute request (%s)", fmt.Sprintf("Response Content-Type is '%s', but should be 'application/json'.", mimeType))
	}

	// Parse Request
	var jsonData *APIResponse

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	// @TODO: handle API exception
	// Argument required
	if jsonData.Code == -2001 {
		return jsonData, fmt.Errorf("Argument error: %s", jsonData.Msg)
	}

	// Permission deny
	if jsonData.Code == -2002 {
		return jsonData, fmt.Errorf("Your point is not enough to use this api")
	}

	return jsonData, nil
}

func (api *TuShare) postData(body map[string]interface{}) (*APIResponse, error) {
	req, err := api.request("POST", Endpoint, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
