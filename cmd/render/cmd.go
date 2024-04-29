package render

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/zhb127/easydev/cmd/render/tmpldir/app"
)

var Cmd = &cli.Command{
	Name:  "render",
	Usage: "渲染生成",
	Subcommands: []*cli.Command{
		{
			Name:        "tmpldir",
			Usage:       "渲染模板目录，生成结果文件",
			Description: "扫描模板目录结构，遍历扫描到的文件列表，识别并渲染模板文件、模板文件名，保存结果文件到输出目录",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "config",
					Aliases:  []string{"c"},
					Usage:    "配置文件（格式：json）",
					Required: true,
				},
				&cli.StringFlag{
					Name:     "input-path",
					Aliases:  []string{"i"},
					Usage:    "输入路径（模板目录路径）",
					Required: false,
				},
				&cli.StringFlag{
					Name:     "output-path",
					Aliases:  []string{"o"},
					Usage:    "输出路径（输出目录路径，不存在将自动创建）",
					Required: false,
				},
				&cli.BoolFlag{
					Name:     "dry-run",
					Usage:    "是否试运行模式（不执行真正操作）",
					Required: false,
				},
				&cli.BoolFlag{
					Name:        "debug",
					Aliases:     []string{"d"},
					Usage:       "是否调试",
					DefaultText: "false",
					Required:    false,
				},
			},
			Action: func(c *cli.Context) error {
				// 解析配置文件
				flagConfig := c.String("config")
				appConfig, err := app.ParseConfig(flagConfig)
				if err != nil {
					return errors.Wrapf(err, "配置文件：%s，解析失败", flagConfig)
				}

				// 合并命令行选项
				flagInputPath := c.String("input-path")
				if flagInputPath != "" {
					appConfig.InputPath = flagInputPath
				}
				flagOutputPath := c.String("output-path")
				if flagOutputPath != "" {
					appConfig.OutputPath = flagOutputPath
				}

				flagDryRun := c.Bool("dry-run")
				if flagDryRun {
					appConfig.DryRun = flagDryRun
				}

				flagDebug := c.Bool("debug")
				if flagDebug {
					appConfig.Debug = flagDebug
				}

				appInst, err := app.New(appConfig)
				if err != nil {
					return err
				}
				return appInst.Run()
			},
		},
	},
}
