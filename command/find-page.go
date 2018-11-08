package command

import (
	"fmt"
	"log"
)

func findPage() {
	validateFindPage()
	result := restClient.SearchPages(options.Title, options.SpaceKey)

	if result.Size > 0 {
		fmt.Println("Page Found!")
		fmt.Println()
	} else {
		fmt.Println("Page not found!")
	}

	for _, element := range result.Results {
		fmt.Println("Title:", element.Title)
		fmt.Println("ID:", element.ID)
		fmt.Println("Version:", element.Version.Number)
	}
}

func validateFindPage() {
	if options.Title == "" || options.SpaceKey == "" {
		log.Fatal("Space Key and Title required to find page!")
	}
}
