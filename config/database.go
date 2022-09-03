package config

//Database 参数
type Database struct {
	Type     string `yaml:"type" json:"type"`
	URL      string `yaml:"url" json:"url"`
	Debug    bool   `yaml:"debug" json:"debug"`
	LogLevel int    `json:"log_level" json:"log_level"`
	Sync     bool   `yaml:"sync" json:"sync"`
}

var DatabaseDefault = Database{
	Type:     "sqlite",
	URL:      "sqlite3.db",
	Debug:    false,
	LogLevel: 4,
	Sync:     true,
}
