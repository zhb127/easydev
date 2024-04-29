package config

import (
	"github.com/pkg/errors"

	"github.com/zhb127/easydev/pkg/tmpl"
)

type TmplVar struct {
	Name       string             `json:"name" validate:"required"`
	Value      string             `json:"value" validate:"omitempty"`
	ValueInput *TmplVarValueInput `json:"value_input" validate:"required_without=Value"`
}

type TmplVarValueInput struct {
	Type string `json:"type" validate:"required,oneof=text select template"`
	Text *struct {
		Label              string `json:"label" validate:"required"`
		DefaultValue       string `json:"default_value"`
		ValueRegexpPattern string `json:"value_regexp_pattern"`
	} `json:"text" validate:"required_if=Type text"`
	Select *struct {
		Label  string   `json:"label" validate:"required"`
		Values []string `json:"values" validate:"required,min=1,dive,required"`
	} `json:"select" validate:"required_if=Type select"`
	Template *struct {
		Text string `json:"text" validate:"required"`
	} `json:"template" validate:"required_if=Type template"`
}

func (c *TmplVarValueInput) ParseToTmplVarValueInputConfig() (*tmpl.VarValueInputConfig, error) {
	typeId, err := tmpl.ParseVarValueInputTypeId(c.Type)
	if err != nil {
		return nil, errors.Errorf("type 为：%s，无效（可选值：%v）", c.Type, tmpl.VarValueInputTypeIds)
	}

	result := &tmpl.VarValueInputConfig{
		Type: typeId,
	}

	switch typeId {
	case tmpl.VarValueInputTypeIdText:
		result.Text = &tmpl.VarValueInputTextConfig{
			Label:              c.Text.Label,
			DefaultValue:       c.Text.DefaultValue,
			ValueRegexpPattern: c.Text.ValueRegexpPattern,
		}
	case tmpl.VarValueInputTypeIdSelect:
		result.Select = &tmpl.VarValueInputSelectConfig{
			Label:  c.Select.Label,
			Values: c.Select.Values,
		}
	case tmpl.VarValueInputTypeIdTemplate:
		result.Template = &tmpl.VarValueInputTemplateConfig{
			Text: c.Template.Text,
		}
	}

	return result, nil
}
