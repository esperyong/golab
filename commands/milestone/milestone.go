package milestone

import (
	"fmt"
	printIssueCmd "github.com/esperyong/golab/commands/milestone/printissue"
	milestoneStatCmd "github.com/esperyong/golab/commands/milestone/stat"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	c := &cli.Command{
		Name:    "milestone",
		Aliases: []string{"m"},
		Usage:   "milestone相关命令",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Milestone")
			return nil
		},
		Subcommands: []*cli.Command{
			milestoneStatCmd.MileStoneStatCommand(),
			printIssueCmd.MileStonePrintIssueCommand(),
		},
	}
	return c
}
