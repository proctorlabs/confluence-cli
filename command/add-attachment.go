package command

import (
	"log"
)

func addAttachment() {
	validateAddAttachment()
	results := restClient.SearchPages(options.Title, options.SpaceKey)
	if results.Size > 0 {
		log.Println("Page found, adding attachment...")
		item := results.Results[0]
		restClient.AddAttachment(options.body, item.ID, options.filename)
		log.Println("Page updated.", item.ID)
	} else {
		log.Fatal("Failed to add attachment, page not found!")
	}
}

func validateAddAttachment() {
	if options.Title == "" || options.SpaceKey == "" || options.Filepath == "" {
		log.Fatal("Space Key, Title, and File Path required for attachment operations!")
	}
}
