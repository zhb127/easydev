package tmpl

import (
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/zhb127/easydev/pkg/spelling"
)

func newRendererKernel() *template.Template {
	return template.New("kernel").Funcs(template.FuncMap{
		// SC
		"ToSnakeCase": strcase.ToSnake, // snake_case
		// SSC
		"ToScreamingSnakeCase": strcase.ToScreamingSnake, // SCREAMING_SNAKE_CASE
		// LCC
		"ToLowerCamelCase": strcase.ToLowerCamel, // lowerCamelCase
		// UCC
		"ToUpperCamelCase": strcase.ToCamel, // UpperCamelCase
		// KC
		"ToKebabCase": strcase.ToKebab, // kebab-case
		// SKC
		"ToScreamingKebabCase": strcase.ToScreamingKebab, // SCREAMING-KEBAB-CASE
		// LC
		"ToLowerCase": strings.ToLower, // lowercase
		// UC
		"ToUpperCase": strings.ToUpper, // UPPERCASE
		// PL
		"ToPlural": spelling.ToPlural, // 复数
		// SG
		"ToSingular": spelling.ToSingular, // 单数
	})
}
