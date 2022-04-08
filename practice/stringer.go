package practice

// -linecomment 使用后面的注释作为描述信息 不加会默认使用定义的常量作为描述
// -output stringer_string.go 生成文件名称 不加会使用 常量类型+string 做为文件名
//go:generate stringer -type=StatusCode -linecomment -output stringer_string.go

type StatusCode int

const (
	Unknown StatusCode = iota //未知
	Success                   //成功
	Failed                    //失败
)