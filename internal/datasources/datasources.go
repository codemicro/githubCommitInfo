package datasources

import (
	"context"
	"time"

	"golang.org/x/oauth2"

	"github.com/shurcooL/githubv4"
)

var readFromYear = 2008

type GithubClient struct {
	gqlClient *githubv4.Client
}

func NewGithubClient(oauthToken string) *GithubClient {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauthToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &GithubClient{
		gqlClient: githubv4.NewClient(tc),
	}

}

func (c GithubClient) GetAllCommits(user string) (int, error) {

	ctx := context.Background()

	// Get commits from GraphQL API

	var query struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions     githubv4.Int
				RestrictedContributionsCount githubv4.Int
				EndedAt                      githubv4.DateTime
			} `graphql:"contributionsCollection(from: $from)"`
		} `graphql:"user(login: $login)"`
	}
	queryVars := map[string]interface{}{
		"login": githubv4.String(user),
	}

	var count int

	stuckReadFrom := readFromYear

	for i := 0; i < time.Now().Year()-stuckReadFrom+1; i++ {

		currentYear := stuckReadFrom + i

		queryVars["from"] = githubv4.DateTime{time.Date(currentYear, time.January, 1, 0, 0, 0, 0, time.UTC)}
		err := c.gqlClient.Query(ctx, &query, queryVars)
		if err != nil {
			return 0, err
		}
		count += int(query.User.ContributionsCollection.TotalCommitContributions)
		count += int(query.User.ContributionsCollection.RestrictedContributionsCount)

		if count == 0 && readFromYear+i >= readFromYear {
			readFromYear = currentYear
		}
	}

	return count, nil
}
