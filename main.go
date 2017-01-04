package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/philproctor/confluence-cli/client"
)

var config = client.ConfluenceConfig{}

var options = client.OperationOptions{}

func main() {
	flag.StringVar(&config.Username, "u", "", "Confluence username")
	flag.StringVar(&config.Password, "p", "", "Confluence password")
	flag.StringVar(&config.URL, "s", "", "The base URL of the Confluence page")
	flag.StringVar(&options.Title, "t", "", "Title to use for a new page")
	flag.StringVar(&options.SpaceKey, "k", "", "Space key to use")
	flag.StringVar(&options.Filepath, "f", "", "Path to the file to upload as the page contents")
	flag.StringVar(&options.AncestorTitle, "A", "", "Title of the ancestor to use")
	flag.Int64Var(&options.AncestorID, "a", 0, "ID of the ancestor to use")
	flag.BoolVar(&config.Debug, "d", false, "Enable debug level logging")
	flag.BoolVar(&options.BodyOnly, "strip-body", false, "If the file is HTML, strip out everything except <body>")
	flag.BoolVar(&options.StripImgs, "strip-imgs", false, "If the file is HTML, strip out all <img> tags")
	command := flag.String("command", "help", "Confluence command to issue")
	flag.Parse()
	runCommand(*command)
}

func runCommand(command string) {
	switch command {
	case "addpage":
		validateBasic()
		validatePageCRUD()
		client.Client(&config).AddOrUpdatePage(options)
		break

	case "searchpage":
		validateBasic()
		result := client.Client(&config).SearchPages(options.Title, options.SpaceKey)
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
	default:
		printUsage()
		break
	}
}

func validateBasic() {
	if config.Password == "" || config.URL == "" || config.Username == "" {
		printUsage()
		log.Fatal("Username, password, and URL required!")
	}
}

func validatePageCRUD() {
	if options.Title == "" || options.SpaceKey == "" || options.Filepath == "" {
		printUsage()
		log.Fatal("Space Key, Title, and File Path required for page operations!")
	}
}

func printUsage() {
	fmt.Println(`
Usage for this Confluence Command Line Interface is as follows:
  -u                  Username to use for Rest API
  -p                  Confluence password to use for Rest API
  -s                  The base URL of the Confluence site
  -a                  Ancestor ID to use for new page
  -A                  Ancestor Title to use for new page
  -t                  The title of the page
  -k                  Space key to use
  -f                  Path to the file for the operation
  -d                  Enable debug level logging
  --strip-body        Strip HTML file to only include contents of <body>
  --strip-imgs        Strip HTML file of all <img> tags
  --command           The command to run against the site
                      Possible values include:
                      addpage: Add a new page to the service
                      searchpage: Search for existing pages that match title
`)
}
