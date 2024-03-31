package handler

// var pagesToScrape []string
// pageToScrape := BaseURL
// pagesDiscovered := []string{pageToScrape}
// fName := "Crypto.xlsx"
// file := excelize.NewFile()

// // Set the sheet name
// err := file.SetSheetName("Sheet1", SheetName)
// if err != nil {
// 	log.Fatal("Not able to Set Sheet Name")
// }

// headers := []string{"Name", "Symbol", "Market Cap (USD)", "Price (USD)"}

// // Write headers to the Excel sheet
// for colIndex, colName := range headers {
// 	file.SetCellValue(SheetName, fmt.Sprintf("%s%d", string(rune(65+colIndex)), 1), colName)
// }

// // current iteration
// i := 1
// // max pages to scrape
// limit := 5
// row := 2 // Start from row 2 for data

// // Instantiate default collector
// c := colly.NewCollector()
// // iterating over the list of pagination links to implement the crawling logic
// c.OnHTML("a.page-numbers", func(e *colly.HTMLElement) {
// 	// discovering a new page
// 	newPaginationLink := e.Attr("href")

// 	// if the page discovered is new
// 	if !contains(pagesToScrape, newPaginationLink) {
// 		// if the page discovered should be scraped
// 		if !contains(pagesDiscovered, newPaginationLink) {
// 			pagesToScrape = append(pagesToScrape, newPaginationLink)
// 		}
// 		pagesDiscovered = append(pagesDiscovered, newPaginationLink)
// 	}
// })
// // Visit the first page and extract data
// c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
// 	// Write data to Excel sheet
// 	file.SetCellValue(SheetName, fmt.Sprintf("A%d", row), e.ChildText(".cmc-table__column-name"))
// 	file.SetCellValue(SheetName, fmt.Sprintf("B%d", row), e.ChildText(".cmc-table__cell--sort-by__symbol"))
// 	file.SetCellValue(SheetName, fmt.Sprintf("C%d", row), e.ChildText(".cmc-table__cell--sort-by__market-cap"))
// 	file.SetCellValue(SheetName, fmt.Sprintf("D%d", row), e.ChildText(".cmc-table__cell--sort-by__price"))
// 	row++
// })

// c.Visit(BaseURL)

// // Initial page scrape
// c.OnScraped(func(response *colly.Response) {
// 	if len(pagesToScrape) != 0 && i < limit {
// 		// getting the current page to scrape and removing it from the list
// 		pageToScrape = pagesToScrape[0]
// 		pagesToScrape = pagesToScrape[1:]

// 		// incrementing the iteration counter
// 		i++

// 		c.Visit(pageToScrape)
// 	}

// })

// // Wait for all scraping to finish
// c.Wait()

// // Save the Excel file
// if err := file.SaveAs(fName); err != nil {
// 	fmt.Println("Error saving Excel file:", err)
// 	return
// }

// log.Printf("Scraping finished, check file %q for results\n", fName)
// }

// func contains(s []string, str string) bool {
// for _, v := range s {
// 	if v == str {
// 		return true
// 	}
// }

// return false
// }
