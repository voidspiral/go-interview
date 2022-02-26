Json中Tag用法汇总
JSON Tag标签格式为：
json:"FieldName/-/可选,omitempty/可选,string/可选
多个选项之间使用 , 逗号分割
FieldName选项：指定编码后键名称
可为空，则使用Struct对应字段名作为JSON输出名
FieldName非空，则使用指定的FieldName作为JSON输出名

- 符号，输出时忽略此字段；但要注意-,（多一个逗号结尾）时，输出字段名为-的JSON字段，而不是忽略
omitempty选项：忽略空值
包含此选项，输出时字段空值（零值+空值：false、0、nil指针、nil接口值，以及任何空数组、切片、map或字符串）则不输出
string选项：结果输出为字符串
字段结果输出为字符串
只适用于字符串、浮点、整数或布尔类型的字段
这种额外的编码有时在与 JavaScript 程序通信时使用
要注意，如果字段值本身为string时，再次增加JSON的string标签选项，会导致多个引号的情况