package client

import (
	"io/ioutil"
	"log"

	"github.com/philproctor/confluence-cli/utility"
)

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, filepath string, bodyOnly, stripImgs bool, ancestor int64) {
	page := newPage(title, spaceKey)
	if ancestor > 0 {
		page.Ancestors = []ConfluencePageAncestor{
			ConfluencePageAncestor{ancestor},
		}
	}
	response := &ConfluencePage{}
	page.Body.Storage.Value = getBodyFromFile(filepath, bodyOnly, stripImgs)
	//page.Body.Storage.Representation = "wiki"
	c.doRequest("POST", "/rest/api/content/", page, response)
	log.Println("ConfluencePage Object Response", response)
}

//UpdatePage adds a new page to the space with the given title
func (c *ConfluenceClient) UpdatePage(title, spaceKey, filepath string, bodyOnly, stripImgs bool, ID string, version, ancestor int64) {
	page := newPage(title, spaceKey)
	page.ID = ID
	page.Version = &ConfluencePageVersion{version}
	if ancestor > 0 {
		page.Ancestors = []ConfluencePageAncestor{
			ConfluencePageAncestor{ancestor},
		}
	}
	response := &ConfluencePage{}
	page.Body.Storage.Value = getBodyFromFile(filepath, bodyOnly, stripImgs)
	//page.Body.Storage.Representation = "wiki"
	c.doRequest("PUT", "/rest/api/content/"+ID, page, response)
	log.Println("ConfluencePage Object Response", response)
}

func getBodyFromFile(filepath string, bodyOnly, stripImgs bool) string {
	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if bodyOnly == false {
		return string(buf)
	}
	return utility.StripHTML(buf, bodyOnly, stripImgs)
}
