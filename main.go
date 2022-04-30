package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/ppark2ya/go-basics/scrapper"
	"github.com/qinains/fastergoding"
)

var FILE_NAME = "jobs.csv"

func handleHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func handleScrape(c *fiber.Ctx) error {
	defer os.Remove(FILE_NAME)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Download(FILE_NAME, FILE_NAME)
}

func main() {
	fastergoding.Run() // hot reload
	engine := html.New("./", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

  app.Get("/", handleHome)
  app.Post("/scrape", handleScrape)

  log.Fatal(app.Listen(":3000"))

}
