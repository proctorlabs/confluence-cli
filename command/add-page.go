package command

import (
	"log"
)

func addPage() {
	validateAddPage()
	addOrUpdatePage()
}

func validateAddPage() {
	if options.Title == "" || options.SpaceKey == "" || options.Filepath == "" {
		log.Fatal("Space Key, Title, and File Path required for page operations!")
	}
}

func addOrUpdatePage() {
	results := restClient.SearchPages(options.Title, options.SpaceKey)
	if results.Size > 0 {
		log.Println("Page found, updating page...")
		item := results.Results[0]
		restClient.UpdatePage(options.Title, options.SpaceKey, options.body, item.ID, item.Version.Number+1, options.AncestorID)
	} else {
		log.Println("Page not found, adding page...")
		restClient.AddPage(options.Title, options.SpaceKey, options.body, options.AncestorID)
	}
}
