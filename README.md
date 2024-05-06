# go-openapi
a tool for building and documenting Go RESTful APIs
基于 github.com/getkin/kin-openapi 下 openapi3 包, 在写法上使用泛型合并了各类操作, 移除了 Unmarshal和Validate相关的代码,仅保留数据结构与 MarshalJSON 和 MarshalYAML, 专注于生成openapi3的文档结构 