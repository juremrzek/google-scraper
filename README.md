# Web scraper for google search

This app allows you to instantly get all the different rankings

Web scraping was done with GoQuery (https://pkg.go.dev/github.com/PuerkitoBio/goquery), the results were formatted from structs into a JSON file using json package (https://pkg.go.dev/encoding/json).
The app includes 2 examples, each made for a specific search query. First query is "sushi", and the second one is "how long to cook pasta"
I used the time package (https://pkg.go.dev/time) to measure performance of the program. It takes around 25ms per 100 results to run (100 results had an average of 27ms, 200 results had 46ms), and it works even if not all result types are present.

To use:
- Go to http://google.com/ and search for a query, then save the HTML.
- Create a json file for selecors.
- Run the following command: `go run main.go`
- Write the name of your HTML file (leaving the field blank will default to sushi.html)
- Write the name of your JSON file (leaving the field blank will default to selectors_sushi.json)