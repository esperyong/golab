package commands

import (
	"github.com/esperyong/golab/commands/issue"
	"github.com/esperyong/golab/commands/milestone"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func ExecCommand() {
	app := &cli.App{
		Name:  "glab",
		Usage: "A GitLab CLI tool bringing base on GitLab, for your XP Dev Scrum Dev.",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "esperyong",
				Email: "liuyong@tuyoo.com",
			},
		},
		Commands: []*cli.Command{
			issue.Command(),
			milestone.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
