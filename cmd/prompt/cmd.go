package prompt

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/zhb127/easydev/errs"
	"github.com/zhb127/easydev/pkg/prompt"
)

var Cmd = &cli.Command{
	Name:  "prompt",
	Usage: "交互式提示",
	Subcommands: []*cli.Command{
		{
			Name:        "select",
			Usage:       "交互式选择",
			Description: "进入交互式选择，根据提示，选择需要的值",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "label",
					Aliases:  []string{"l"},
					Usage:    "提示内容",
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
					if errors.Is(err, prompt.ErrInterrupt) {
						return errs.ErrInterrupt
					}
					return err
				}

				result := resultInAny.(string)

				// 将结果输出到标准输出，以便在 SHELL 中能使用 $() 捕获结果
				fmt.Print(result)

				return nil
			},
		},
		{
			Name:        "text",
			Usage:       "交互式文本输入",
			Description: "进入交互式文本输入，按提示，输入符合格式的文本",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "label",
					Aliases:  []string{"l"},
					Usage:    "提示内容",
					Required: true,
				},
				&cli.StringSliceFlag{
					Name:     "value-regexp-pattern",
					Aliases:  []string{"v"},
					Usage:    "输入值格式（正则表达式模式）",
					Required: false,
				},
			},
			Action: func(c *cli.Context) error {
				flagLabel := c.String("label")
				flagValueRegexpPattern := c.String("value-regexp-pattern")

				appInst, err := prompt.NewTextPrompt(&prompt.TextPromptConfig{
					Label:              flagLabel,
					ValueRegexpPattern: flagValueRegexpPattern,
					Stdout:             os.Stderr, // 将生成的交互界面输出到标准错误
				})
				if err != nil {
					return err
				}

				resultInAny, err := appInst.Run()
				if err != nil {
					if errors.Is(err, prompt.ErrInterrupt) {
						return errs.ErrInterrupt
					}
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
