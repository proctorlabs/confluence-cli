# Confluence CLI

This is a command line interface for Confluence. Usage of the command line is as below:

```
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
```

As such, some example commands that can be run:

To add or update a page with the title "New Page Title".
```
confluence-cli -u test-user -p test-password -s http://localhost:8080/wiki --command addpage -k TST -t "New Page Title" -f path/to/file
```

Same as above, accept ensure the page is underneath "Ancestor Page" in the wiki. Use -a instead to add underneath by ID instead of title
```
confluence-cli -u test-user -p test-password -s http://localhost:8080/wiki --command addpage -k TST -A "Ancestor Page" -t "New Page Title" -f path/to/file
```
