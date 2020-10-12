package command

import (
	"log"

	"github.com/proctorlabs/confluence-cli/client"
)

var restClient *client.ConfluenceClient

//Run the provided command with the options from the command line
func Run(command string, config *client.ConfluenceConfig, opts *OperationOptions) {

	options = opts

	switch command {
	case "add-page":
		prepareClient(config)
		addPage()

	case "add-or-update-page":
		prepareClient(config)
		addOrUpdatePage()

	case "update-page":
		prepareClient(config)
		updatePage()

	case "find-page":
		prepareClient(config)
		findPage()

	case "add-attachment":
		prepareClient(config)
		addAttachment()

	case "help", "":
		printUsage()

	default:
		log.Fatal("Unknown Command: ", command)
	}

}

func prepareClient(config *client.ConfluenceConfig) {
	validateClientDetails(config)
	restClient = client.Client(config)
	processFlags()
}
