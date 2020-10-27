package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codemicro/githubCommitInfo/internal/endpoints"
	"github.com/codemicro/githubCommitInfo/internal/shields"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	oauthToken := os.Getenv("GITHUB_OAUTH_TOKEN")
	userName := os.Getenv("GITHUB_USERNAME")

	if oauthToken == "" {
		fmt.Println("GITHUB_OAUTH_TOKEN envirnoment variable not set.")
		os.Exit(1)
	}

	if userName == "" {
		fmt.Println("GITHUB_USERNAME environment variable not set.")
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(shields.NewShield(c.Locals("fieldName").(string), "Unavailable", "red"))
		},
	})

	app.Get("/", cache.New(), endpoints.NewCommitEndpoint(userName, oauthToken))

	log.Panic(app.Listen(":80"))
}
