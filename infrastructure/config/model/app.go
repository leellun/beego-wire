package model

type Config struct {
	App   *App   `yaml:"app"`
	Log   *Log   `yaml:"log"`
	Pgsql *Pgsql `yaml:"pgsql"`
}

type App struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Log struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}

type Pgsql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	SslMode  string `yaml:"sslMode"`
}
