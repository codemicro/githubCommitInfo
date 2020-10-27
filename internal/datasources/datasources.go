package datasources

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/shurcooL/githubv4"
)

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

	var count int

	// Get commits from GraphQL API

	var query struct {
		User struct {
			ContributionsCollection struct {
				TotalCommitContributions     githubv4.Int
				RestrictedContributionsCount githubv4.Int
			}
		} `graphql:"user(login: $login)"`
	}
	queryVars := map[string]interface{}{
		"login": githubv4.String(user),
	}

	err := c.gqlClient.Query(ctx, &query, queryVars)
	if err != nil {
		return count, err
	}

	count = int(query.User.ContributionsCollection.TotalCommitContributions) + int(query.User.ContributionsCollection.RestrictedContributionsCount)

	return count, nil
}
