package client

import (
	"fmt"
	"io/ioutil"
)

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, filepath string) {
	page := newPage(title, spaceKey)
	response := &ConfluencePage{}
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	page.Body.Storage.Value = string(buf)
	c.doRequest("POST", "/rest/api/content/", page, response)
	fmt.Println("ConfluencePage Object Response", response)
}
