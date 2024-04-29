//go:generate go run github.com/abice/go-enum --marshal
package tmpl

// 变量值输入器类型 ID
/*
ENUM(
text		// 文本
select		// 选择
template	// 模板
)
*/
type VarValueInputTypeId string

var VarValueInputTypeIds = []VarValueInputTypeId{
	VarValueInputTypeIdText,
	VarValueInputTypeIdSelect,
	VarValueInputTypeIdTemplate,
}
