package stat

import (
	"fmt"
	"github.com/esperyong/golab/glab"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"log"
	"time"
)

func milestoneStat(groupId int, mileStoneName string) {
	git, err := glab.GitlabClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	issues, _, err := git.Issues.ListGroupIssues(groupId,
		&gitlab.ListGroupIssuesOptions{
			Milestone: &mileStoneName,
			ListOptions: gitlab.ListOptions{
				Page:    1,
				PerPage: 1000,
			}})
	if err != nil {
		log.Fatalf("Failed to list issue: %v", err)
	}
	issueNum := len(issues)
	if issueNum == 0 {
		log.Fatalln("没找到相关的Issue，有可能groupId或mileStone名称输错了")
	} else {
		seconds := 0
		for _, issue := range issues {
			seconds += issue.TimeStats.TimeEstimate
		}
		mileStoneEstimateSum := time.Second * time.Duration(seconds)
		storyPoint := mileStoneEstimateSum.Hours() / 8
		s := fmt.Sprintf(
			"MileStone:[%s]\n共:[%d]个Issue\n预估总开发时间:[%s]\n故事点：[%.1f]个\n",
			mileStoneName, issueNum, mileStoneEstimateSum, storyPoint)
		fmt.Println(s)
	}
}

func MileStoneStatCommand() *cli.Command {
	mileStoneName := ""
	groupId := 5
	c := &cli.Command{
		Name:  "stat",
		Usage: "statistic milestone",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mname",
				Aliases:     []string{"m"},
				Value:       mileStoneName,
				Required:    true,
				Usage:       "milestone名称",
				Destination: &mileStoneName,
			},
			&cli.IntFlag{
				Name:        "group",
				Aliases:     []string{"g"},
				Value:       groupId,
				Required:    true,
				Usage:       "Group ID",
				Destination: &groupId,
			},
		},
		Action: func(cCtx *cli.Context) error {
			milestoneStat(groupId, mileStoneName)
			return nil
		},
	}
	return c
}
