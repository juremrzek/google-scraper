# Web scraper for google search

This app allows you to get different rankings of google search results (organic, local, carousel, knowledge panel, and featured snippet)

Web scraping was done with GoQuery (https://pkg.go.dev/github.com/PuerkitoBio/goquery), the results were formatted from structs into a JSON file using json package (https://pkg.go.dev/encoding/json).
The app includes 2 examples, each made for a specific search query. First query is "sushi", and the second one is "how long to cook pasta".

To use:
- Go to http://google.com/ and search for a query, then save the HTML.
- Create a json file for selecors.
- Run the following command: `go run main.go` and pass html and json files as arguments
- Example: `go run main.go search_results.html selectors.json`

