package client

import (
	"log"
)

//AddPage adds a new page to the space with the given title
func (c *ConfluenceClient) AddPage(title, spaceKey, body string, ancestor int64) {
	page := newPage(title, spaceKey)
	if ancestor > 0 {
		page.Ancestors = []ConfluencePageAncestor{
			ConfluencePageAncestor{ancestor},
		}
	}
	response := &ConfluencePage{}
	page.Body.Storage.Value = body
	//page.Body.Storage.Representation = "wiki"
	c.doRequest("POST", "/rest/api/content/", page, response)
	log.Println("ConfluencePage Object Response", response)
}

//UpdatePage adds a new page to the space with the given title
func (c *ConfluenceClient) UpdatePage(title, spaceKey, body string, ID string, version, ancestor int64) {
	page := newPage(title, spaceKey)
	page.ID = ID
	page.Version = &ConfluencePageVersion{version}
	if ancestor > 0 {
		page.Ancestors = []ConfluencePageAncestor{
			ConfluencePageAncestor{ancestor},
		}
	}
	response := &ConfluencePage{}
	page.Body.Storage.Value = body
	//page.Body.Storage.Representation = "wiki"
	c.doRequest("PUT", "/rest/api/content/"+ID, page, response)
	log.Println("ConfluencePage Object Response", response)
}
