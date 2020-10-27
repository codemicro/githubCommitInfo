package main

import (
	"fmt"

	"github.com/codemicro/githubCommitInfo/internal/datasources"
)

func main() {
	client := datasources.NewClient("e86aa6438a9f39efd1485340d19426ea24eda6a9")
	num, err := client.GetAllCommits("codemicro")
	if err != nil {
		panic(err)
	}
	fmt.Println(*num)
}
