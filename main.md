package main

import (
	"os"
	"strings"

	"github.com/honor1111/Nomad/scrapper"
	"github.com/labstack/echo"
)

const fileNAME string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileNAME)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileNAME, fileNAME)
}
func handleHome(c echo.Context) error {
	return c.File("home.html")
}
