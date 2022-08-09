package printissue

import (
	"fmt"
	"github.com/esperyong/golab/glab"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/urfave/cli/v2"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
	"strings"
	"time"
)

// WordWrap 在指定的列数处换行文本
// s string 字符串文本数据
// limit int 限制单词个数
func WordWrap(sentence string, limit int) string {
	if strings.TrimSpace(sentence) == "" {
		return sentence
	}
	//将字符串转换为切片
	var result string = ""
	var i = 0
	for _, word := range sentence {
		i++
		if i%limit == 0 {
			result = result + fmt.Sprintf("%c\n", word)
		} else {
			result = result + fmt.Sprintf("%c", word)
		}

	}
	return result
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func printIssues(groupId int, mileStoneName string, fileType string) pdf.Maroto {
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
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	for i, issue := range issues {
		buildIssueCard(m, issue)
		if i%2 == 1 {
			m.AddPage()
		}
	}
	return m
}

//解决中文TODO(目前中文换行还是有问题)
//https://github.com/johnfercher/maroto/issues/142
//https://github.com/Vale-sail/maroto
//https://github.com/rongfengliang/maroto
func buildIssueCard(m pdf.Maroto, issue *gitlab.Issue) {
	m.SetBorder(true)
	rowHeight := 5.0
	m.AddUTF8Font("NotoSansSC", "", "./fonts/NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "I", "./fonts/NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "B", "./fonts/NotoSansSC-Regular.ttf")
	m.AddUTF8Font("NotoSansSC", "BI", "./fonts/NotoSansSC-Regular.ttf")
	textProps := props.Text{
		Top:    0.5,
		Family: "NotoSansSC",
		Style:  consts.Bold,
		Align:  consts.Left,
	}

	m.Row(rowHeight*1, func() {
		m.Col(12, func() {
			m.Text(fmt.Sprintf("故事ID: %s", issue.WebURL), textProps)
		})
	})
	m.Row(rowHeight*2, func() {
		m.Col(2, func() {
			m.Text("故事名称:", textProps)
		})
		m.Col(10, func() {
			m.Text(issue.Title, props.Text{
				Size:   13,
				Family: "NotoSansSC",
				Align:  consts.Left,
			})
		})
	})
	m.Row(rowHeight*1, func() {
		m.Col(6, func() {
			m.Text(fmt.Sprintf("重要性: %d", issue.Weight), textProps)
		})
		m.Col(6, func() {
			seconds := issue.TimeStats.TimeEstimate
			mileStoneEstimateSum := time.Second * time.Duration(seconds)
			storyPoint := mileStoneEstimateSum.Hours() / 8
			s := fmt.Sprintf("故事点：[%.1f]个", storyPoint)
			m.Text(s, textProps)
		})
	})

	m.Row(90, func() {
		m.Col(12, func() {
			text := fmt.Sprintf("故事描述：%s", issue.Description)
			m.Text(text, props.Text{
				Extrapolate: false,
				Size:        12,
				Top:         2,
				Family:      "NotoSansSC",
				Align:       consts.Left,
			})
		})
	})

	m.SetBorder(false)

	m.Row(rowHeight*1, func() {
		m.Col(12, func() {
		})
	})

}

func MileStonePrintIssueCommand() *cli.Command {
	mileStoneName := ""
	fileType := "pdf"
	groupId := 5
	c := &cli.Command{
		Name:    "printissue",
		Usage:   "print issues in a milestone for Print Physical Scrum Board Story Card",
		Aliases: []string{"pi"},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "filetype",
				Aliases:     []string{"t"},
				Value:       fileType,
				Required:    false,
				Usage:       "输出格式",
				Destination: &fileType,
			},
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
			m := printIssues(groupId, mileStoneName, fileType)
			err := m.OutputFileAndClose(fmt.Sprintf("./%s.pdf", mileStoneName))
			if err != nil {
				fmt.Println("⚠️  Could not save PDF:", err)
				os.Exit(1)
			}
			fmt.Println("PDF saved successfully")
			return nil
		},
	}
	return c
}
