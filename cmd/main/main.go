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
	"etf-data-extractor/api/routes"

	"github.com/gin-gonic/gin"
)

const (
	SheetName = "CryptoData"
	BaseURL   = "https://coinmarketcap.com/all/views/all/"
)

func main() {
	router := gin.Default()

	routes.Register(router)
	router.Run("localhost:8080")
}
