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
  -s                  Confluence site base url (e.g. https://example.atlassian.net/wiki)

command flags
  -a                  Ancestor ID to use for new page
  -A                  Ancestor Title to use for new page
  -t                  The title of the page
  -k                  Space key to use
  -f                  Path to the file to process/upload
  -R                  Representation of the file to upload (storage, wiki, can be any supported by confluence convert api)
  -d                  Enable debug level logging
  --strip-body        Strip HTML file to only include contents of <body>
  --strip-imgs        Strip HTML file of all <img> tags
  --clean-adoc        Aggressively cleans HTML generated from .adoc to make it play nicely with confluence

  <command>           The command to run
                         add-page: Add a new page to Confluence
                         add-or-update-page: Add a new page to Confluence or update if it already exists
                         update-page: Update an existing page on confluence
                         add-attachment: Add or update an attachment to the specified page
                         find-page: Search for existing pages that match title`)
}
