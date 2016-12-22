package client

import (
	"fmt"
	"io/ioutil"

	"github.com/philproctor/confluence-cli/utility"
)

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, filepath string, bodyOnly, stripImgs bool) {
	page := newPage(title, spaceKey)
	response := &ConfluencePage{}
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	if bodyOnly == false {
		page.Body.Storage.Value = string(buf)
	} else {
		page.Body.Storage.Value = utility.StripHTML(buf, bodyOnly, stripImgs)
	}
	c.doRequest("POST", "/rest/api/content/", page, response)
	fmt.Println("ConfluencePage Object Response", response)
}
