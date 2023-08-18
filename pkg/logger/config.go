package logger

type Zap struct {
	Level         string `yaml:"level"`                                        // 级别
	Prefix        string `yaml:"prefix"`                                       // 日志前缀
	Format        string `yaml:"format"`                                       // 输出
	Director      string `yaml:"director"`                                     // 日志文件夹
	EncodeLevel   string `yaml:"encode-level" mapstructure:"encode-level"`     // 编码级
	StacktraceKey string `yaml:"stacktrace-key" mapstructure:"stacktrace-key"` // 栈名
	MaxAge        int    `yaml:"max-age" mapstructure:"max-age"`               // 日志留存时间
	ShowLine      bool   `yaml:"show-line" mapstructure:"show-line"`           // 显示行
	LogInConsole  bool   `yaml:"log-in-console" mapstructure:"log-in-console"` // 输出控制台
}
