package command

import (
	"fmt"
	"log"
)

func findPage() {
	validateFindPage()
	result := restClient.SearchPages(options.Title, options.SpaceKey)
	fmt.Println("Pages Found: ", result.Size)
	fmt.Println()
	for index, element := range result.Results {
		fmt.Println("Page", index)
		fmt.Println("Title:", element.Title)
		fmt.Println("ID:", element.ID)
		fmt.Println("Type:", element.Type)
		fmt.Println("Version:", element.Version.Number)
		fmt.Println()
	}
}

func validateFindPage() {
	if options.Title == "" || options.SpaceKey == "" {
		log.Fatal("Space Key and Title required to find page!")
	}
}
