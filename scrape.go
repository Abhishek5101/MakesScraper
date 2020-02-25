package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
	"net/http"
)

type news struct {
	Title []string `json: "title"`
}

func startServer(myData []string) {
	e := echo.New()

	e.GET("/", func(f echo.Context) error {
		return f.JSON(http.StatusOK, myData)
	})

	fmt.Println("Server running: http://localhost:8000")
	e.Logger.Fatal(e.Start(":8000"))
}

func main() {
	var myData [] string

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a.storylink", func(e *colly.HTMLElement) {
		myData = append(myData, e.Text)
		title := news{Title:myData}
		byteArray, err := json.MarshalIndent(title, "", " ")
		if err != nil{
			fmt.Println(err)
		}
		file, err := os.Create("output.json")
		if err != nil{
			fmt.Print(err)
		}
		file.Write(byteArray)
		defer file.Close()
		fmt.Println(string(byteArray))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://news.ycombinator.com
	c.Visit("https://news.ycombinator.com/")
	startServer(myData)
}
