package client

//ConvertToStorage returns the supplied text in storage format
func (c *ConfluenceClient) ConvertToStorage(body, sourceFormat string) string {
	convertRequest := &ConfluenceConvert{
		Value:          body,
		Representation: sourceFormat,
	}
	c.doRequest("POST", "/rest/api/contentbody/convert/storage", convertRequest, convertRequest)
	return convertRequest.Value
}
