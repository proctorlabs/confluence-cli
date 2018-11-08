# Confluence CLI

This is a command line interface for Confluence. Usage of the command line is as below:

```
Usage for this Confluence Command Line Interface is as follows:
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
                         find-page: Search for existing pages that match title
```

As such, some example commands that can be run:

To add or update a page with the title "New Page Title".
```
confluence-cli -u test-user -p test-password -s http://localhost:8080/wiki -k TST -t "New Page Title" -f path/to/file add-page
```

Same as above, except ensure the page is underneath "Ancestor Page" in the wiki. Use -a instead to add underneath by ID instead of title
```
confluence-cli -u test-user -p test-password -s http://localhost:8080/wiki -k TST -A "Ancestor Page" -t "New Page Title" -f path/to/file add-page
```
