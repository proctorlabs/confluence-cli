package command

import (
	"log"

	"github.com/philproctor/confluence-cli/client"
)

func addPage(config *client.ConfluenceConfig, options *client.OperationOptions) {
	validateClientDetails(config)
	validateAddPage(options)
	client.Client(config).AddOrUpdatePage(options)
}

func validateAddPage(options *client.OperationOptions) {
	if options.Title == "" || options.SpaceKey == "" || options.Filepath == "" {
		printUsage()
		log.Fatal("Space Key, Title, and File Path required for page operations!")
	}
}
