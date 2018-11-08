package command

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/philproctor/confluence-cli/utility"
)

//OperationOptions holds all the options that apply to the specified operation
type OperationOptions struct {
	Title         string
	SpaceKey      string
	Filepath      string
	BodyOnly      bool
	StripImgs     bool
	AncestorTitle string
	AncestorID    int64
	body          string
}

var options *OperationOptions

func processFlags() {
	if options.AncestorTitle != "" {
		setAncestorFromTitle()
	}
	if options.Filepath != "" {
		processAndSetBody()
	}
}

func setAncestorFromTitle() {
	ancestorResults := restClient.SearchPages(options.AncestorTitle, options.SpaceKey)
	if ancestorResults.Size < 1 {
		log.Fatal("Ancestor title not found!")
	} else {
		ancestorIDint, err := strconv.Atoi(ancestorResults.Results[0].ID)
		log.Println("Found ancestor ID", ancestorIDint)
		if err != nil {
			log.Fatal(err)
		}
		options.AncestorID = int64(ancestorIDint)
	}
}

func processAndSetBody() {
	buf, err := ioutil.ReadFile(options.Filepath)
	if err != nil {
		log.Fatal(err)
	}
	if options.BodyOnly == false && options.StripImgs == false {
		options.body = string(buf)
	} else {
		options.body = utility.StripHTML(buf, options.BodyOnly, options.StripImgs)
	}
	log.Println("Successfully processed file: ", options.Filepath)
}
