# Confluence CLI

This is a command line interface for Confluence. Usage of the command line is as below:

``` bash
Usage for this Confluence Command Line Interface is as follows:
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
                         find-page: Search for existing pages that match title
```

As such, some example commands that can be run:

To add a page with the title "New Page Title".

``` bash
confluence-cli -u test-user -p test-password -s https://example.atlassian.net/wiki -k TST -t "New Page Title" -f path/to/file add-page
```

Same as above, except ensure the page is underneath "Ancestor Page" in the wiki. Use -a instead to add underneath by ID instead of title

``` bash
confluence-cli -u test-user -p test-password -s https://example.atlassian.net/wiki -k TST -A "Ancestor Page" -t "New Page Title" -f path/to/file add-page
```

The easiest way to install and run this tool is with docker. Note that you must volume mount the file you want to upload inside the container.

``` bash
docker run --rm -v /path/to/file.txt:/in.txt philproctor/confluence-cli -u test-user -p test-password -s https://example.atlassian.net/wiki -A 'Ancestor Page' -k TST -f /in.txt -t 'Page Title' -R 'wiki' add-or-update-page
```
