package client

//AddAttachment adds an attachment to an existing page
func (c *ConfluenceClient) AddAttachment(content, pageID, filename string) {
	results := &ConfluencePageSearch{}
	c.uploadFile("PUT", "/rest/api/content/"+pageID+"/child/attachment", content, filename, &results)
}
