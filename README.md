# easydev

一款简单的开发辅助命令行工具。

包含：

- render
  - tmpldir：渲染模板目录输出
- prompt
  - select：交互式选择

## 渲染模板目录输出

easydev render tmpldir

选项：

- -i, -input-dir：模板目录
- -o, -output-dir：输出目录
- -c, -config：配置文件
- -d, -debug：调试（输出 DEBUG 日志）
- -dry-run：试运行模式（不执行真实操作）

### 模板目录示例

```c
// 文件位置：test/data/render-tmpldir/template

└── template                                // 模板目录
    ├── biz                                 // 待创建的目录（已存在，则跳过）
    │   └── {{.ResourceNameSC}}.go.tmpl     // 待生成的文件（文件名、文件内容：模板解析；文件名：移除模板文件后缀名）
    ├── dao                                 // 待创建的目录
    │   └── mysql                           // 待创建的目录
    │       └── {{.ResourceNameSC}}.go.tmpl // 待生成的文件（已存在，则跳过）
    └── tmpl_vars.txt.tmpl                  // 待生成的文件（文件内容：模板解析；文件名：移除模板文件后缀名）
```

> 示例文件路径：test/data/render-tmpldir/template

### 配置文件示例

> 文件位置：configs/render-tmpldir.json

```c
{
    "tmpl_file_ext": ".tmpl",               // 模板文件后缀名
    "tmpl_vars": [
        {
            "name": "ResourceNameSC",                                       // 模板变量名
            "value_input": {                                                // 模板变量值输入器配置
                "type": "text",                                             // 输入器类型：提示文本输入
                "text": {                                                   // 输入器配置
                    "label": "资源标识名（蛇形命名，如：api_token）",            // - 提示文案
                    "value_regexp_pattern": "^[a-z]{2,}(_[a-z]{2,}){0,8}$"  // - 输入值的正则表达式模式
                }
            }
        },
        {
            "name": "DatabaseDriver",
            "value_input": {
                "type": "select",           // 输入器类型：提示文本选择
                "select": {                 // 输入器配置
                    "label": "数据库驱动",    // - 提示文案
                    "values": [             // - 可选值数组
                        "MySQL",
                        "ClickHouse",
                        "PostgreSQL"
                    ]
                }
            }
        },
        {
            "name": "DatabaseDriverLC",
            "value_input": {
                "type": "template",                                 // 输入器类型：模板解析
                "template": {                                       // 输入器配置
                    "text": "{{ .DatabaseDriver | ToLowerCase }}"   // - 模板文本（仅支持 go template 语法，用于将之前输入的模板变量值再处理）
                }
            }
        },
        {
            "name": "ResourceNameUCC",
            "value_input": {
                "type": "template",
                "template": {
                    "text": "{{ .ResourceNameSC | ToUpperCamelCase }}"
                }
            }
        },
        {
            "name": "ResourceNameConst",
            "value": "abc"                  // 无输入器，固定值
        }
    ]
}
```
