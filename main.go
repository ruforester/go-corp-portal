package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// router := gin.Default()
	e := echo.New()
	// router.Static("/", "./static")
	e.Static("/", "static")
	// router.LoadHTMLGlob("templates/*.html")
	e.File("/", "templates/axentix.html")

	e.GET("/links", func(c echo.Context) error {
		return c.HTML(http.StatusOK, csvToHtml("links.csv", "a", ','))
	})

	e.GET("/reqs", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h2></h2>")
	})

	e.GET("/dir", func(c echo.Context) error {
		return c.HTML(http.StatusOK, csvToHtml("example2.csv", "table", ','))
	})

	print(csvToHtml("example2.csv", "table", ','))

	e.Logger.Fatal(e.Start(":1323"))

	// router.GET("/dir", func(ctx *gin.Context) {
	// 	ctx.Data(http.StatusOK, "text/html", csvToHtml("example.csv"))
	// })

	// router.GET("/links", func(ctx *gin.Context) {
	// 	ctx.Data(http.StatusOK, "text/html", mdToHtml("links.md"))
	// })

	// router.Run(":8080")

}

func csvToHtml(csvPath, htmlElement string, sep rune) string {
	f, err := os.Open(csvPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = sep
	r.LazyQuotes = true
	s, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	if htmlElement == "table" {
		// htmlTable := "<table class='table table-striped table-responsive'>" //bootstrap
		htmlTable := "<table class='table table-striped table-scroll'>" //spectre
		for i, row := range s {
			if i > 20 {
				break
			}
			if i == 0 {
				htmlTable = htmlTable + "<thead>"
			}
			htmlTable = htmlTable + "<tr>"
			for _, col := range row {
				if i == 0 {
					htmlTable = htmlTable + fmt.Sprintf("<th>%s</th>", col)
				} else {
					htmlTable = htmlTable + fmt.Sprintf("<td>%s</td>", col)
				}
			}
			htmlTable = htmlTable + "</tr>"
			if i == 0 {
				htmlTable = htmlTable + "</thead>"
			}
		}
		htmlTable = htmlTable + "</table>"
		return htmlTable
	}

	href := "<ul>"
	for i, row := range s {
		if i == 0 {
			continue
		}
		href = href + fmt.Sprintf("<li><a href='%s'>%s</a></li>", row[1], row[0])
	}
	href = href + "</ul>"
	return href
}

// func mdToHtml(mdPath string) []byte {
// 	f, _ := os.Open(mdPath)
// 	b, _ := io.ReadAll(f)
// 	mdBytes := blackfriday.Run(b)
// 	return mdBytes
// }
