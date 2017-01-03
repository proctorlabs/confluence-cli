package client

import (
	"fmt"
	"io/ioutil"

	"github.com/philproctor/confluence-cli/utility"
)

//AddOrUpdatePage checks for an existing page then calls AddPage or UpdatePage depending on the result
func (c *ConfluenceClient) AddOrUpdatePage(title, spaceKey, filepath string, bodyOnly, stripImgs bool) {
	results := c.SearchPages(title, spaceKey)
	if results.Size > 0 {
		fmt.Println("Page found, updating page...")
		item := results.Results[0]
		c.UpdatePage(title, spaceKey, filepath, bodyOnly, stripImgs, item.ID, item.Version.Number+1)
	} else {
		fmt.Println("Page not found, adding page...")
		c.AddPage(title, spaceKey, filepath, bodyOnly, stripImgs)
	}
}

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, filepath string, bodyOnly, stripImgs bool) {
	page := newPage(title, spaceKey)
	response := &ConfluencePage{}
	page.Body.Storage.Value = getBodyFromFile(filepath, bodyOnly, stripImgs)
	c.doRequest("POST", "/rest/api/content/", page, response)
	fmt.Println("ConfluencePage Object Response", response)
}

//UpdatePage adds a new page to the space with the given title
func (c *ConfluenceClient) UpdatePage(title, spaceKey, filepath string, bodyOnly, stripImgs bool, ID string, version int64) {
	page := newPage(title, spaceKey)
	page.ID = ID
	page.Version = &ConfluencePageVersion{version}
	response := &ConfluencePage{}
	page.Body.Storage.Value = getBodyFromFile(filepath, bodyOnly, stripImgs)
	c.doRequest("PUT", "/rest/api/content/"+ID, page, response)
	fmt.Println("ConfluencePage Object Response", response)
}

func getBodyFromFile(filepath string, bodyOnly, stripImgs bool) string {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	if bodyOnly == false {
		return string(buf)
	}
	return utility.StripHTML(buf, bodyOnly, stripImgs)
}
