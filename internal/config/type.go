package config

type Config struct {
	Server   Server   `yaml:"server"`
	Database DBConfig `yaml:"database"`
}

type Server struct {
	Host         string `yaml:"host"`
	Port         int32  `yaml:"port"`
	WriteTimeout int64  `yaml:"write_timeout"`
	ReadTimeout  int64  `yaml:"read_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	DBName   string `yaml:"db_name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
