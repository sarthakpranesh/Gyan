package main

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type Gyan struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

func main() {
	app := fiber.New()
	memCache := cache.New(24*time.Hour, 1*time.Hour)

	// Server Info Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Gyan")
	})

	// Scrape Data Route
	app.Get("/:name", func(c *fiber.Ctx) error {
		var info Gyan
		var name string = c.Params("name")
		tmp, wasThere := memCache.Get(name)
		if wasThere {
			info = tmp.(Gyan)
			return c.JSON(info)
		}

		info.Name = name

		scrape := colly.NewCollector()

		scrape.OnHTML(".mw-parser-output", func(h *colly.HTMLElement) {
			var description string
			h.ForEachWithBreak("p", func(i int, h *colly.HTMLElement) bool {
				if i > 3 {
					return false
				}
				p := h.Text
				// clean the text
				p = strings.ReplaceAll(p, "\n", "")
				reg := regexp.MustCompile(`\[[0-9]{1,}\]`)
				p = reg.ReplaceAllString(p, "")
				description = description + p
				return true
			})
			info.Description = description
		})

		var link string = "https://en.wikipedia.org/wiki/" + name
		info.Link = link
		scrape.Visit(link)
		scrape.Wait()
		memCache.Add(name, info, 24*time.Hour)
		return c.JSON(info)
	})

	app.Listen(":8080")
}
