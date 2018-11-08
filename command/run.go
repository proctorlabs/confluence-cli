package command

import (
	"github.com/philproctor/confluence-cli/client"
)

//Run the provided command with the options from the command line
func Run(command string, config *client.ConfluenceConfig, options *client.OperationOptions) {

	switch command {
	case "add-page":
		addPage(config, options)
		break

	case "find-page":
		findPage(config, options)
		break

	default:
		printUsage()
		break
	}
}
