package issue

import (
	"fmt"
	issuePrintCmd "github.com/esperyong/golab/commands/issue/print"
	"github.com/urfave/cli/v2"
	"os"
)

func Command() *cli.Command {
	c := &cli.Command{
		Name:    "issue",
		Aliases: []string{"i"},
		Usage:   "issue相关命令",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Issue \n")
			token := os.Getenv("GITLAB_TOKEN")
			url := os.Getenv("GITLAB_BASE_URL")
			fmt.Println("token with url:", token, url)
			return nil
		},
		Subcommands: []*cli.Command{issuePrintCmd.IssuePrintCommand()},
	}
	return c
}
