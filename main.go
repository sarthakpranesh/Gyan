package main

import (
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type Gyan struct {
	Name        string   `json:"name"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
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
		info.Name = c.Params("name")
		if info.Name == "" {
			return c.JSON(info)
		}
		tmp, wasThere := memCache.Get(info.Name)
		if wasThere {
			info = tmp.(Gyan)
			return c.JSON(info)
		}

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			scrape := colly.NewCollector()
			scrape.OnHTML("img", func(h *colly.HTMLElement) {
				imageLink := h.Attr("src")
				if strings.Contains(imageLink, "https") {
					info.Images = append(info.Images, imageLink)
				}
			})
			var imageLink string = "https://www.google.com/search?tbm=isch&q=" + info.Name
			scrape.Visit(imageLink)
			scrape.Wait()
			wg.Done()
		}()
		go func() {
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
			var link string = "https://en.wikipedia.org/wiki/" + info.Name
			info.Link = link
			scrape.Visit(link)
			scrape.Wait()
			wg.Done()
		}()

		// Wait for all the go routines to finish
		wg.Wait()

		// Cache scraped content
		memCache.Add(info.Name, info, 24*time.Hour)
		return c.JSON(info)
	})

	app.Listen(":8080")
}
