package app

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
	"github.com/zhb127/easydev/cmd/render/tmpldir/app/config"
	"github.com/zhb127/easydev/pkg/file"
	"github.com/zhb127/easydev/pkg/validate"
)

type Config struct {
	TmplFileExt string            `json:"tmpl_file_ext" validate:"required"`
	TmplVars    []*config.TmplVar `json:"tmpl_vars" validate:"omitempty,dive,required"`
	InputPath   string            `json:"input_path" validate:"required"`
	OutputPath  string            `json:"output_path" validate:"required"`
	DryRun      bool              `json:"dry_run"`
	Debug       bool              `json:"debug"`
}

func ParseConfig(filePath string) (*Config, error) {
	fileContentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "读取文件内容失败")
	}

	result := &Config{}
	if err := json.Unmarshal(fileContentBytes, result); err != nil {
		return result, errors.Wrapf(err, "按 JSON 解析文件内容失败")
	}

	return result, nil
}

func validateConfig(config *Config) error {
	if config == nil {
		return errors.New("config 不能为空")
	}

	if err := validate.Struct(config); err != nil {
		return err
	}

	if !file.IsDir(config.InputPath) {
		return errors.Errorf("input_path=%s 不是有效的目录", config.InputPath)
	}

	return nil
}
