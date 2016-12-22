package main

import (
	"flag"
	"fmt"

	"github.com/philproctor/confluence-cli/client"
)

var config = client.ConfluenceConfig{}

var options cliOptions

type cliOptions struct {
	title     string
	spaceKey  string
	filepath  string
	bodyOnly  bool
	stripImgs bool
}

func main() {
	flag.StringVar(&config.Username, "u", "", "Confluence username")
	flag.StringVar(&config.Password, "p", "", "Confluence password")
	flag.StringVar(&config.URL, "s", "", "The base URL of the Confluence page")
	flag.StringVar(&options.title, "t", "", "Title to use for a new page")
	flag.StringVar(&options.spaceKey, "k", "", "Space key to use")
	flag.StringVar(&options.filepath, "f", "", "Path to the file to upload as the page contents")
	flag.BoolVar(&options.bodyOnly, "strip-body", false, "If the file is HTML, strip out everything except <body>")
	flag.BoolVar(&options.stripImgs, "strip-imgs", false, "If the file is HTML, strip out all <img> tags")
	command := flag.String("command", "help", "Confluence command to issue")
	flag.Parse()
	runCommand(*command)
}

func runCommand(command string) {
	switch command {
	case "addpage":
		validateBasic()
		validatePageCRUD()
		client.Client(&config).AddPage(options.title, options.spaceKey, options.filepath, options.bodyOnly, options.stripImgs)
		break

	default:
		printUsage()
		break
	}
}

func validateBasic() {
	if config.Password == "" || config.URL == "" || config.Username == "" {
		printUsage()
		panic("Username, password, and URL required!")
	}
}

func validatePageCRUD() {
	if options.title == "" || options.spaceKey == "" || options.filepath == "" {
		printUsage()
		panic("Space Key, Title, and File Path required for page operations!")
	}
}

func printUsage() {
	fmt.Println(`
Usage for this Confluence Command Line Interface is as follows:
  -u                  Username to use for Rest API
  -p                  Confluence password to use for Rest API
  -s                  The base URL of the Confluence site
  -t                  The title of the page
  -k                  Space key to use
  -f                  Path to the file for the operation
  --strip-body        Strip HTML file to only include contents of <body>
  --strip-imgs        Strip HTML file of all <img> tags
  --command           The command to run against the site
                      Possible values include:
                      addpage: Add a new page to the service
`)
}
