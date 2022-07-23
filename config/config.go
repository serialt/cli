package config

import "time"

type Service struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

type Database struct {
	Type            string        `yaml:"type"`
	DBName          string        `gorm:"dbname" yaml:"dbname"`
	Addr            string        `gorm:"addr" yaml:"addr"`
	Port            string        `gorm:"port" yaml:"port"`
	Username        string        `gorm:"username" yaml:"username"`
	Password        string        `gorm:"password" yaml:"password"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}
type Log struct {
	LogLevel string `yaml:"logLevel"` // 日志级别，支持debug,info,warn,error,panic
	LogFile  string `yaml:"logFile"`  // 日志文件存放路径,如果为空，则输出到控制台
}

type MyConfig struct {
	Log      Log      `json:"log" yaml:"log"`
	Service  Service  `json:"service" yaml:"service"`
	Database Database `json:"database" yaml:"database"`
}
