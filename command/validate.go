package command

import (
	"fmt"
	"log"

	"github.com/philproctor/confluence-cli/client"
)

func validateClientDetails(config *client.ConfluenceConfig) {
	if config.Password == "" || config.URL == "" || config.Username == "" {
		log.Fatal("Username, password, and URL required for this operation!")
	}
}

func printUsage() {
	fmt.Println(`Usage for this Confluence Command Line Interface is as follows:
confluence-cli [flags] <command>

authentication
  -u                  Confluence username
  -p                  Confluence password
  -s                  Confluence site base url (e.g. https://companyname.atlassian.net/wiki)

command flags
  -a                  Ancestor ID to use for new page
  -A                  Ancestor Title to use for new page
  -t                  The title of the page
  -k                  Space key to use
  -f                  Path to the file to process/upload
  -d                  Enable debug level logging
  --strip-body        Strip HTML file to only include contents of <body>
  --strip-imgs        Strip HTML file of all <img> tags

  <command>           The command to run
                         add-page: Add a new page to the service
                         find-page: Search for existing pages that match title`)
}
