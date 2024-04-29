package tmpl

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/zhb127/easydev/pkg/prompt"
	"github.com/zhb127/easydev/pkg/validate"
)

type VarValueInputTextConfig struct {
	Label              string `validate:"required"`
	DefaultValue       string
	ValueRegexpPattern string
}

type VarValueInputSelectConfig struct {
	Label  string   `validate:"required"`
	Values []string `validate:"required,gt=0,dive,required"`
}

type VarValueInputTemplateConfig struct {
	Text      string `validate:"required"`
	VarValues map[string]interface{}
}

type VarValueInputConfig struct {
	Type     VarValueInputTypeId
	Text     *VarValueInputTextConfig     `validate:"required_if=Type text"`
	Select   *VarValueInputSelectConfig   `validate:"required_if=Type select"`
	Template *VarValueInputTemplateConfig `validate:"required_if=Type template"`
}

type VarValueInput struct {
	config *VarValueInputConfig
}

func NewVarValueInput(config *VarValueInputConfig) (*VarValueInput, error) {
	if config == nil {
		return nil, errors.New("配置不能为空")
	}
	if err := validate.Struct(config); err != nil {
		return nil, errors.Wrap(err, "配置验证失败")
	}

	return &VarValueInput{
		config,
	}, nil
}

func (i *VarValueInput) Run() (interface{}, error) {
	switch i.config.Type {
	case VarValueInputTypeIdText, VarValueInputTypeIdSelect:
		p, err := func() (prompt.Prompt, error) {
			switch i.config.Type {
			case VarValueInputTypeIdText:
				return prompt.NewTextPrompt(&prompt.TextPromptConfig{
					Label:              i.config.Text.Label,
					DefaultValue:       i.config.Text.DefaultValue,
					ValueRegexpPattern: i.config.Text.ValueRegexpPattern,
					Stdout:             os.Stderr,
				})
			case VarValueInputTypeIdSelect:
				return prompt.NewSelectPrompt(&prompt.SelectPromptConfig{
					Label:  i.config.Select.Label,
					Values: i.config.Select.Values,
					Stdout: os.Stderr,
				})
			}
			return nil, errors.Errorf("类型：%s，不支持", i.config.Type)
		}()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for {
			result, err := p.Run()
			if err != nil {
				if errors.Is(err, prompt.ErrInterrupt) {
					return nil, ErrVarValueInputInterrupt
				}
				fmt.Println(err)
				continue
			}
			return result, nil
		}
	case VarValueInputTypeIdTemplate:
		tmplRenderer := NewRenderer(nil)
		result, err := tmplRenderer.RenderTmplText(i.config.Template.Text, i.config.Template.VarValues)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return result, nil
	default:
		return nil, errors.Errorf("类型：%s，不支持", i.config.Type)
	}
}
