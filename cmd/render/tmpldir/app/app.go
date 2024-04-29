package app

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/zhb127/easydev/pkg/file"
	"github.com/zhb127/easydev/pkg/log"
	"github.com/zhb127/easydev/pkg/tmpl"
)

type app struct {
	config *Config

	fileScanner   *file.Scanner
	tmplRenderer  *tmpl.Renderer
	fileGenerator *file.Generator
}

func New(config *Config) (*app, error) {
	if err := validateConfig(config); err != nil {
		return nil, errors.Wrap(err, "配置验证失败")
	}

	if config.Debug {
		log.SetDebugMode()
	}

	fileScanner := file.NewScanner()
	tmplRenderer := tmpl.NewRenderer(fileScanner)
	fileGenerator := file.NewGenerator(fileScanner)

	return &app{
		config,
		fileScanner,
		tmplRenderer,
		fileGenerator,
	}, nil
}

func (a *app) Run() error {
	log.SetPrefix("[构建模板变量]")
	tmplVarValues, err := a.buildTmplVarValues()
	if err != nil {
		if errors.Is(err, ErrInterrupt) {
			return nil
		}
		return errors.WithStack(err)
	}

	log.SetPrefix("[渲染模板目录]")
	fileInfosRendered, err := a.tmplRenderer.RenderTmplDir(&tmpl.RenderTmplDirParams{
		TmplVarValues: tmplVarValues,
		TmplFileExt:   a.config.TmplFileExt,
		TmplDir:       a.config.InputPath,
		OutputDir:     a.config.OutputPath,
	})
	if err != nil {
		log.Err(err)
		return errors.WithStack(err)
	}

	log.SetPrefix("[生成输出文件]")
	fileInfosGenerated, err := a.fileGenerator.GenFilesByFileInfos(fileInfosRendered, a.config.DryRun)
	if err != nil {
		return errors.WithStack(err)
	}
	log.SetPrefix("")

	for _, v := range fileInfosGenerated {
		fmt.Println(v.Path)
	}

	return nil
}

// 构建模板变量值
func (a *app) buildTmplVarValues() (map[string]interface{}, error) {
	log.Debugf("开始")

	result := make(map[string]interface{})

	for k, v := range a.config.TmplVars {
		// 固定值
		if v.Value != "" {
			result[v.Name] = v.Value
			continue
		}

		tmplVarValueInputConfig, err := v.ValueInput.ParseToTmplVarValueInputConfig()
		if err != nil {
			return nil, errors.Wrapf(err, "配置项：tmpl_vars[%d]，解析 value_input 失败", k)
		}

		if tmplVarValueInputConfig.Type == tmpl.VarValueInputTypeIdTemplate {
			tmplVarValueInputConfig.Template.VarValues = result
		}

		tmplVarValueInput, err := tmpl.NewVarValueInput(tmplVarValueInputConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "配置项：tmpl_vars[%d]，TmplVarValueInput 实例化失败", k)
		}

		tmplVarValue, err := tmplVarValueInput.Run()
		if err != nil {
			if errors.Is(err, tmpl.ErrVarValueInputInterrupt) {
				return nil, ErrInterrupt
			}
			return nil, errors.Wrapf(err, "配置项：tmpl_vars[%d]，TmplVarValueInput 运行失败", k)
		}

		result[v.Name] = tmplVarValue
	}

	return result, nil
}
