package command

import (
	"fmt"
	"log"

	"github.com/philproctor/confluence-cli/client"
)

func findPage(config *client.ConfluenceConfig, options *client.OperationOptions) {
	validateClientDetails(config)
	result := client.Client(config).SearchPages(options.Title, options.SpaceKey)
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

func validateFindPage(options *client.OperationOptions) {
	if options.Title == "" || options.SpaceKey == "" {
		printUsage()
		log.Fatal("Space Key and Title required to find page!")
	}
}
