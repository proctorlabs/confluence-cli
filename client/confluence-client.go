package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"time"
)

//ConfluenceClient is the primary client to the Confluence API
type ConfluenceClient struct {
	username string
	password string
	baseURL  string
	debug    bool
	client   *http.Client
}

//Client returns a new instance of the client
func Client(config *ConfluenceConfig) *ConfluenceClient {
	return &ConfluenceClient{
		username: config.Username,
		password: config.Password,
		baseURL:  config.URL,
		debug:    config.Debug,
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *ConfluenceClient) doRequest(method, url string, content, responseContainer interface{}) []byte {
	b := new(bytes.Buffer)
	if content != nil {
		json.NewEncoder(b).Encode(content)
	}
	furl := c.baseURL + url
	if c.debug {
		log.Println("Full URL", furl)
		log.Println("JSON Content:", b.String())
	}
	request, err := http.NewRequest(method, furl, b)
	request.SetBasicAuth(c.username, c.password)
	request.Header.Add("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		log.Fatal(err)
	}
	if c.debug {
		log.Println("Sending request to services...")
	}
	response, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if c.debug {
		log.Println("Response received, processing response...")
		log.Println("Response status code is", response.StatusCode)
		log.Println(response.Status)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if c.debug {
		log.Println("Response from service...", string(contents))
	}
	if response.StatusCode != 200 {
		log.Fatal("Bad response code received from server: ", response.Status)
	}
	json.Unmarshal(contents, responseContainer)
	return contents
}

func (c *ConfluenceClient) uploadFile(method, url, content, filename string, responseContainer interface{}) []byte {
	b := new(bytes.Buffer)
	writer := multipart.NewWriter(b)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		log.Fatal(err)
	}
	part.Write([]byte(content))
	writer.WriteField("minorEdit", "true")
	//writer.WriteField("comment", "test")
	writer.Close()

	furl := c.baseURL + url
	if c.debug {
		log.Println("Full URL", furl)
	}
	request, err := http.NewRequest(method, furl, b)
	request.SetBasicAuth(c.username, c.password)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("X-Atlassian-Token", "nocheck")
	if err != nil {
		log.Fatal(err)
	}
	if c.debug {
		log.Println("Sending request to services...")
	}
	response, err := c.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if c.debug {
		log.Println("Response received, processing response...")
		log.Println("Response status code is", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		log.Fatal("Bad response code received from server: ", response.Status)
	}
	json.Unmarshal(contents, responseContainer)
	return contents
}
