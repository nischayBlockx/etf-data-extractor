package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

const (
	SheetName = "CryptoData"
)

func FetchCryptoData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		// Instantiate default collector
		c := colly.NewCollector()
		row := 2 // Start from row 2 for data

		// Visit the website and extract data
		c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
			// Write data to Excel sheet
			file.SetCellValue(SheetName, fmt.Sprintf("A%d", row), e.ChildText(".cmc-table__column-name"))
			file.SetCellValue(SheetName, fmt.Sprintf("B%d", row), e.ChildText(".cmc-table__cell--sort-by__symbol"))
			file.SetCellValue(SheetName, fmt.Sprintf("C%d", row), e.ChildText(".cmc-table__cell--sort-by__market-cap"))
			file.SetCellValue(SheetName, fmt.Sprintf("D%d", row), e.ChildText(".cmc-table__cell--sort-by__price"))
			file.SetCellValue(SheetName, fmt.Sprintf("E%d", row), e.ChildText(".cmc-table__cell--sort-by__circulating-supply"))
			file.SetCellValue(SheetName, fmt.Sprintf("F%d", row), e.ChildText(".cmc-table__cell--sort-by__volume-24-h"))
			file.SetCellValue(SheetName, fmt.Sprintf("G%d", row), e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"))
			file.SetCellValue(SheetName, fmt.Sprintf("H%d", row), e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"))
			file.SetCellValue(SheetName, fmt.Sprintf("I%d", row), e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"))

			// Increment the row counter
			row++
		})

		// Visit the website
		c.Visit("https://coinmarketcap.com/all/views/all/")

		// Save the Excel file
		if err := file.SaveAs(fName); err != nil {
			fmt.Println("Error saving Excel file:", err)
		}

		log.Printf("Scraping finished, check file %q for results\n", fName)
	}

}
