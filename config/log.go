package config

type LogOutput struct {
	Filename   string `yaml:"filename" json:"filename"`
	MaxSize    int    `yaml:"max_size" json:"max_size,omitempty"`
	MaxAge     int    `yaml:"max_age" json:"max_age,omitempty"`
	MaxBackups int    `yaml:"max_backups" json:"max_backups,omitempty"`
}

//Log 参数
type Log struct {
	Level  string    `yaml:"level" json:"level"`
	Text   bool      `yaml:"text" json:"text,omitempty"`
	Format string    `yaml:"format,omitempty" json:"format,omitempty"`
	Output LogOutput `yaml:"output" json:"output"`
}

var LogDefault = Log{
	Level: "trace",
	Text:  false,
}
