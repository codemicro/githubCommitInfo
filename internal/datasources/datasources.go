package datasources

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v32/github"
)

type GithubClient struct {
	client *github.Client
}

func NewClient(oauthToken string) *GithubClient {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauthToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &GithubClient{
		client: github.NewClient(tc),
	}

}

func (c GithubClient) GetAllCommits(user string) (*int, error) {
	result, _, err := c.client.Search.Commits(context.Background(), fmt.Sprintf("author:%s ", user), nil)
	if err != nil {
		var zero int
		return &zero, err
	}

	return result.Total, nil
}
