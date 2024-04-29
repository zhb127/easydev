package prompt

import (
	"fmt"
	"io"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/zhb127/easydev/pkg/validate"
)

type TextPromptConfig struct {
	Label              string `validate:"required"`
	DefaultValue       string
	ValueRegexpPattern string
	Stdout             io.WriteCloser
}

type textPrompt struct {
	kernel *promptui.Prompt
}

func NewTextPrompt(config *TextPromptConfig) (Prompt, error) {
	if config == nil {
		return nil, errors.New("配置不能为空")
	}

	if err := validate.Struct(config); err != nil {
		return nil, errors.Wrap(err, "配置验证失败")
	}

	label := config.Label
	valueRegexpPattern := config.ValueRegexpPattern
	defaultValue := config.DefaultValue

	var valueRegexp *regexp.Regexp
	if valueRegexpPattern != "" {
		if tmp, err := regexp.Compile(valueRegexpPattern); err != nil {
			return nil, errors.Wrap(err, "value_regexp_pattern 编译错误")
		} else {
			valueRegexp = tmp
		}
	}

	if valueRegexpPattern != "" {
		label += fmt.Sprintf("（格式：%s）", valueRegexpPattern)
	}

	kernel := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	if config.Stdout != nil {
		kernel.Stdout = config.Stdout
	}

	if valueRegexp != nil {
		kernel.Validate = func(v string) error {
			if matched := valueRegexp.MatchString(v); !matched {
				return errors.Errorf("输入值：%s，不符合格式：%s", v, valueRegexpPattern)
			}
			return nil
		}
	}

	p := &textPrompt{
		kernel: &kernel,
	}

	return p, nil
}

func (p *textPrompt) Run() (interface{}, error) {
	result, err := p.kernel.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrInterrupt) {
			return nil, ErrInterrupt
		}
		return nil, err
	}
	return result, nil
}
