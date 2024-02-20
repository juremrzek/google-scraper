package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type JSONElementIn struct {
	RankType    string `json:"rankType"`
	Selector    string `json:"selector"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type Information struct {
	Rank        int    `json:"rank"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

type JSONElementOut struct {
	RankType string        `json:"rankType"`
	Info     []Information `json:"info"`
}

func main() {

	fmt.Print("Enter the name of html file: ")
	var input string
	fmt.Scanln(&input)
	if len(input) == 0 {
		input = "sushi.html"
	}
	fmt.Println(input)

	htmlFile, err := os.ReadFile(input) //Here is the input html file
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlFile))
	if err != nil {
		panic(err)
	}
	input = ""
	fmt.Print("Enter the name of json file: ")
	fmt.Scanln(&input)
	if len(input) == 0 {
		input = "selectors_sushi.json"
	}

	jsonFile, err := os.ReadFile(input) //Here is the input json file
	if err != nil {
		panic(err)
	}
	fmt.Println(input)

	var jsonElements []JSONElementIn

	//
	if err := json.Unmarshal(jsonFile, &jsonElements); err != nil {
		panic(err)
	}

	file, err := os.Create("result.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var dataArray []JSONElementOut

	for _, el := range jsonElements {

		var infoArray []Information

		doc.Find(el.Selector).Each(func(rank int, s *goquery.Selection) {

			url, _ := s.Find(el.Url).First().Attr("href")
			title := s.Find(el.Title).First().Text()
			desc := s.Find(el.Description).First().Text()

			info := Information{
				Rank:        rank + 1,
				Url:         url,
				Title:       title,
				Description: desc,
			}
			infoArray = append(infoArray, info)

		})

		data := JSONElementOut{
			RankType: el.RankType,
			Info:     infoArray,
		}
		dataArray = append(dataArray, data)

	}

	jsonData, err := json.MarshalIndent(dataArray, "", " ")
	if err != nil {
		panic(err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON file created.")

}
