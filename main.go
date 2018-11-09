package main

import (
	"flag"

	"github.com/philproctor/confluence-cli/client"
	"github.com/philproctor/confluence-cli/command"
)

func main() {
	var config = client.ConfluenceConfig{}
	var options = command.OperationOptions{}

	flag.StringVar(&config.Username, "u", "", "Confluence username")
	flag.StringVar(&config.Password, "p", "", "Confluence password")
	flag.StringVar(&config.URL, "s", "", "The base URL of the Confluence page")
	flag.StringVar(&options.Title, "t", "", "Title to use for a new page")
	flag.StringVar(&options.SpaceKey, "k", "", "Space key to use")
	flag.StringVar(&options.Filepath, "f", "", "Path to the file to upload as the page contents")
	flag.StringVar(&options.Format, "R", "storage", "Representation of the file to be uploaded")
	flag.StringVar(&options.AncestorTitle, "A", "", "Title of the ancestor to use")
	flag.Int64Var(&options.AncestorID, "a", 0, "ID of the ancestor to use")
	flag.BoolVar(&config.Debug, "d", false, "Enable debug level logging")
	flag.BoolVar(&options.BodyOnly, "strip-body", false, "If the file is HTML, strip out everything except <body>")
	flag.BoolVar(&options.StripImgs, "strip-imgs", false, "If the file is HTML, strip out all <img> tags")
	flag.Parse()

	command.Run(flag.Arg(0), &config, &options)
}
