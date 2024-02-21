package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type CSSSelector struct {
	RankType    string `json:"rankType"`
	Result      string `json:"result"` //selector of the entire search result element.
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type SearchResultInformation struct {
	Rank        int    `json:"rank"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type SearchResult struct {
	RankType string                    `json:"rankType"`
	Info     []SearchResultInformation `json:"info"`
}

func main() {

	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	} else {
		input = "sushi.html"
	}

	htmlFile, err := os.ReadFile(input) //Read the input html file
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile))
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 2 {
		input = os.Args[2]
	} else {
		input = "selectors_sushi.json"
	}

	jsonFile, err := os.ReadFile(input) //Read the input json file
	if err != nil {
		panic(err)
	}
	fmt.Println(input)

	var cssSelectors []CSSSelector

	//parse the input JSON file
	if err := json.Unmarshal(jsonFile, &cssSelectors); err != nil {
		panic(err)
	}

	searchResults := getSearchResultInformation(doc, cssSelectors)

	//Parse search result data into JSON data
	jsonData, err := json.MarshalIndent(searchResults, "", " ")
	if err != nil {
		panic(err)
	}

	//Write JSON data into a file
	file, err := os.Create("result.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON file created.")

}

func getSearchResultInformation(doc *goquery.Document, selectors []CSSSelector) []SearchResult {
	var searchResults []SearchResult

	//Loop over each rank type
	for _, el := range selectors {

		var infoArray []SearchResultInformation

		//Loop over instances of the rank type
		doc.Find(el.Result).Each(func(rank int, s *goquery.Selection) {

			url, _ := s.Find(el.Url).First().Attr("href")
			title := s.Find(el.Title).First().Text()
			desc := s.Find(el.Description).First().Text()

			info := SearchResultInformation{
				Rank:        rank + 1,
				Url:         url,
				Title:       title,
				Description: desc,
			}
			infoArray = append(infoArray, info)

		})

		searchResult := SearchResult{
			RankType: el.RankType,
			Info:     infoArray,
		}
		searchResults = append(searchResults, searchResult)
	}
	return searchResults
}
