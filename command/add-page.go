package command

import (
	"log"
	"strconv"
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
	ancestorID := options.AncestorID
	if options.AncestorTitle != "" {
		ancestorResults := restClient.SearchPages(options.AncestorTitle, options.SpaceKey)
		if ancestorResults.Size < 1 {
			log.Fatal("Ancestor title not found!")
		} else {
			ancestorIDint, err := strconv.Atoi(ancestorResults.Results[0].ID)
			log.Println("Found ancestor ID", ancestorIDint)
			if err != nil {
				log.Fatal(err)
			}
			ancestorID = int64(ancestorIDint)
		}
	}
	if results.Size > 0 {
		log.Println("Page found, updating page...")
		item := results.Results[0]
		restClient.UpdatePage(options.Title, options.SpaceKey, options.Filepath, options.BodyOnly, options.StripImgs, item.ID, item.Version.Number+1, ancestorID)
	} else {
		log.Println("Page not found, adding page...")
		restClient.AddPage(options.Title, options.SpaceKey, options.Filepath, options.BodyOnly, options.StripImgs, ancestorID)
	}
}
