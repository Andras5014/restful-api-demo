package conf

const (
	TextFormat = LogFormat("text")
	JSONFormat = LogFormat("json")
)

// LogFormat 日志格式
type LogFormat string

const (
	// FileLogTo 输出到文件
	FileLogTo = LogTo("file")
	// StdoutLogTo 输出到控制台
	StdoutLogTo = LogTo("stdout")
)

// LogTo 日志记录地方
type LogTo string
