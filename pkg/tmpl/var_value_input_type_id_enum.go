// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package tmpl

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	// VarValueInputTypeIdText is a VarValueInputTypeId of type text.
	// 文本
	VarValueInputTypeIdText VarValueInputTypeId = "text"
	// VarValueInputTypeIdSelect is a VarValueInputTypeId of type select.
	// 选择
	VarValueInputTypeIdSelect VarValueInputTypeId = "select"
	// VarValueInputTypeIdTemplate is a VarValueInputTypeId of type template.
	// 模板
	VarValueInputTypeIdTemplate VarValueInputTypeId = "template"
)

var ErrInvalidVarValueInputTypeId = errors.New("not a valid VarValueInputTypeId")

// String implements the Stringer interface.
func (x VarValueInputTypeId) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x VarValueInputTypeId) IsValid() bool {
	_, err := ParseVarValueInputTypeId(string(x))
	return err == nil
}

var _VarValueInputTypeIdValue = map[string]VarValueInputTypeId{
	"text":     VarValueInputTypeIdText,
	"select":   VarValueInputTypeIdSelect,
	"template": VarValueInputTypeIdTemplate,
}

// ParseVarValueInputTypeId attempts to convert a string to a VarValueInputTypeId.
func ParseVarValueInputTypeId(name string) (VarValueInputTypeId, error) {
	if x, ok := _VarValueInputTypeIdValue[name]; ok {
		return x, nil
	}
	return VarValueInputTypeId(""), fmt.Errorf("%s is %w", name, ErrInvalidVarValueInputTypeId)
}

// MarshalText implements the text marshaller method.
func (x VarValueInputTypeId) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *VarValueInputTypeId) UnmarshalText(text []byte) error {
	tmp, err := ParseVarValueInputTypeId(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
