package config

//Web 参数
type Web struct {
	Addr     string `yaml:"addr" json:"addr"`
	Debug    bool   `yaml:"debug" json:"debug"`
	Compress bool   `yaml:"compress" json:"compress"`
}

var WebDefault = Web{
	Addr:     ":8080",
	Debug:    true,
	Compress: true,
}
