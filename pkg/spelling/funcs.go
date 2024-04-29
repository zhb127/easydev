package spelling

import (
	pluralize "github.com/gertd/go-pluralize"
)

var pluralizeClient = pluralize.NewClient()

// 转为复数
func ToPlural(word string) string {
	return pluralizeClient.Plural(word)
}

// 转为单数
func ToSingular(word string) string {
	return pluralizeClient.Singular(word)
}
