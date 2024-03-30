// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// func main() {
// 	c := colly.NewCollector(
// 		// Allow visiting only the flipkart.com domain
// 		colly.AllowedDomains("www.flipkart.com"),
// 	)

// 	c.WithRequestOptions(colly.RequestOptions{
// 		Timeout: 60 * time.Second,
// 	})
// 	// Slice to store mobile data
// 	var mobiles []map[string]string

// 	// Visit the Flipkart mobiles page
// 	c.OnHTML("div._1AtVbE", func(e *colly.HTMLElement) {
// 		mobile := make(map[string]string)

// 		// Extract mobile name
// 		mobile["name"] = e.ChildText("a._1fQZEK")

// 		// Extract mobile price
// 		mobile["price"] = e.ChildText("div._30jeq3")

// 		// Extract mobile description
// 		mobile["description"] = e.ChildText("li.tVe95H")

// 		// Append mobile data to the slice
// 		mobiles = append(mobiles, mobile)
// 	})

// 	// Error handling
// 	c.OnError(func(r *colly.Response, err error) {
// 		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})

// 	// Start scraping
// 	c.Visit("https://www.flipkart.com/mobiles/pr?sid=tyy,4io&otracker=categorytree")

//		// Print the fetched mobile data
//		for _, mobile := range mobiles {
//			fmt.Println("Name:", mobile["name"])
//			fmt.Println("Price:", mobile["price"])
//			fmt.Println("Description:", mobile["description"])
//			fmt.Println()
//		}
//	}
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

const (
	SheetName = "CryptoData"
	BaseURL   = "https://coinmarketcap.com/all/views/all/"
)

func main() {
	fName := "Crypto.xlsx"
	file := excelize.NewFile()

	// Set the sheet name
	err := file.SetSheetName("Sheet1", SheetName)
	if err != nil {
		log.Fatal("Not able to Set Sheet Name")
	}

	headers := []string{"Name", "Symbol", "Market Cap (USD)", "Price (USD)"}

	// Write headers to the Excel sheet
	for colIndex, colName := range headers {
		file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+colIndex)), 1), colName)
	}

	row := 2 // Start from row 2 for data

	// Instantiate default collector
	c := colly.NewCollector()

	// Visit the first page and extract data
	visitPage := func(url string) {
		c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
			// Write data to Excel sheet
			file.SetCellValue(SheetName, fmt.Sprintf("A%d", row), e.ChildText(".cmc-table__column-name"))
			file.SetCellValue(SheetName, fmt.Sprintf("B%d", row), e.ChildText(".cmc-table__cell--sort-by__symbol"))
			file.SetCellValue(SheetName, fmt.Sprintf("C%d", row), e.ChildText(".cmc-table__cell--sort-by__market-cap"))
			file.SetCellValue(SheetName, fmt.Sprintf("D%d", row), e.ChildText(".cmc-table__cell--sort-by__price"))
			row++
		})

		c.Visit(url)
	}

	// Initial page scrape
	visitPage(BaseURL)

	// Callback for pagination links
	c.OnHTML(".pagination-container > .paginationBottom > .paginations a", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/all/views/all/") {
			visitPage("https://coinmarketcap.com" + link)
		}
	})

	// Wait for all scraping to finish
	c.Wait()

	// Save the Excel file
	if err := file.SaveAs(fName); err != nil {
		fmt.Println("Error saving Excel file:", err)
		return
	}

	log.Printf("Scraping finished, check file %q for results\n", fName)
}
