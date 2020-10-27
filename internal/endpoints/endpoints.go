package endpoints

import (
	"fmt"
	"strconv"

	"github.com/codemicro/githubCommitInfo/internal/datasources"
	"github.com/codemicro/githubCommitInfo/internal/shields"
	"github.com/gofiber/fiber/v2"
)

func NewCommitEndpoint(username, oauthToken string) fiber.Handler {
	client := datasources.NewGithubClient(oauthToken)
	fieldName := fmt.Sprintf("%s's total commits", username)

	return func(c *fiber.Ctx) error {
		numCommits, err := client.GetAllCommits(username)
		if err != nil {
			c.Locals("fieldName", fieldName)
			return err
		}
		return c.JSON(shields.NewShield(fieldName, strconv.Itoa(numCommits), "green"))
	}
}
