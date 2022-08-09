package glab

import (
	"github.com/xanzy/go-gitlab"
	"os"
)

func GitlabClient() (*gitlab.Client, error) {
	token := os.Getenv("GITLAB_TOKEN")
	url := os.Getenv("GITLAB_BASE_URL")
	git, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	return git, err
}
