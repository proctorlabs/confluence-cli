package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//ConfluenceClient is the primary client to the Confluence API
type ConfluenceClient struct {
	username string
	password string
	baseURL  string
	client   *http.Client
}

//Client returns a new instance of the client
func Client(config *ConfluenceConfig) *ConfluenceClient {
	return &ConfluenceClient{
		username: config.Username,
		password: config.Password,
		baseURL:  config.URL,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *ConfluenceClient) doRequest(method, url string, content, responseContainer interface{}) []byte {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(content)
	request, err := http.NewRequest(method, c.baseURL+url, b)
	request.SetBasicAuth(c.username, c.password)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		panic(err)
	}
	fmt.Println("Sending request to services...")
	response, err := c.client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response received, processing response...")
	fmt.Println("Response status code is", response.StatusCode)
	fmt.Println(response.Status)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		fmt.Println("Bad response code received from server!")
		panic(string(contents))
	}
	json.Unmarshal(contents, responseContainer)
	return contents
}
