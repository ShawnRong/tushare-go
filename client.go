package TuShare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
)

// API endpoint
const API_URL = "http://api.tushare.pro"

type TuShare struct {
	token  string
	client *http.Client
}

func New(token string) *TuShare {
	return NewWithClient(token, http.DefaultClient)
}

func NewWithClient(token string, httpClient *http.Client) *TuShare {
	return &TuShare{
		token:  token,
		client: httpClient,
	}
}

func (api *TuShare) request(method, path string, body interface{}) (*http.Request, error) {
	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (api *TuShare) doRequest(req *http.Request) (*ApiResponse, error) {
	// Set http content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := api.client.Do(req)
	//Handle network error
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Oops! Network error.")
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
	var jsonData *ApiResponse

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

func (api *TuShare) postData(body map[string]interface{}) (*ApiResponse, error) {
	req, err := api.request("POST", API_URL, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
