package prompt

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zhb127/easydev/pkg/prompt"
)

var Cmd = &cli.Command{
	Name:  "prompt",
	Usage: "交互式提示",
	Subcommands: []*cli.Command{
		{
			Name:        "select",
			Usage:       "交互式选择",
			Description: "进入交互式选择，按操作说明，输入一个值",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "label",
					Aliases:  []string{"l"},
					Usage:    "提示文案",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     "values",
					Aliases:  []string{"v"},
					Usage:    "可选值列表（csv 格式）",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				flagLabel := c.String("label")
				flagValues := c.StringSlice("values")

				appInst, err := prompt.NewSelectPrompt(&prompt.SelectPromptConfig{
					Label:  flagLabel,
					Values: flagValues,
					Stdout: os.Stderr, // 将生成的交互界面输出到标准错误
				})
				if err != nil {
					return err
				}

				resultInAny, err := appInst.Run()
				if err != nil {
					return err
				}

				result := resultInAny.(string)

				// 将结果输出到标准输出，以便在 SHELL 中能使用 $() 捕获结果
				fmt.Print(result)

				return nil
			},
		},
	},
}
