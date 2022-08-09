package print

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func IssuePrintCommand() *cli.Command {
	c := &cli.Command{
		Name:  "print",
		Usage: "print issues to pdf",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("print issues: ", cCtx.Args().First())
			return nil
		},
	}
	return c
}
