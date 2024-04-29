package prompt

import (
	"io"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/zhb127/easydev/pkg/validate"
)

type SelectPromptConfig struct {
	Label  string   `validate:"required"`
	Values []string `validate:"required,min=1,dive,required"`
	Stdout io.WriteCloser
}

type selectPrompt struct {
	kernel *promptui.Select
}

func NewSelectPrompt(config *SelectPromptConfig) (Prompt, error) {
	if config == nil {
		return nil, errors.New("配置不能为空")
	}

	if err := validate.Struct(config); err != nil {
		return nil, errors.Wrap(err, "配置验证失败")
	}

	valuesForSearch := make([]string, 0, len(config.Values))
	for _, v := range config.Values {
		valueForSearch := strings.ToLower(v)
		valuesForSearch = append(valuesForSearch, valueForSearch)
	}

	kernel := &promptui.Select{
		Label:             config.Label,
		Items:             config.Values,
		StartInSearchMode: true,
		Searcher: func(inputText string, idx int) bool {
			lowerInputText := strings.ToLower(inputText)
			return strings.Contains(valuesForSearch[idx], lowerInputText)
		},
	}

	if config.Stdout != nil {
		kernel.Stdout = config.Stdout
	}

	result := &selectPrompt{
		kernel,
	}

	return result, nil
}

func (p *selectPrompt) Run() (interface{}, error) {
	_, selectedValue, err := p.kernel.Run()
	if err != nil {
		if errors.Is(err, promptui.ErrInterrupt) {
			return nil, ErrInterrupt
		}
		return nil, err
	}

	return selectedValue, nil
}
