package config

import (
	"go.uber.org/zap/zapcore"
)

type Config struct {
	System *systemConfig `mapstructure:"system" json:"system"`
	Mysql  *mysqlConfig  `mapstructure:"mysql" json:"mysql"`
	Logs   *logsConfig   `mapstructure:"logs" json:"logs"`
	Redis  *redisConfig  `mapstructure:"redis" json:"redis"`
}

type systemConfig struct {
	Mode            string `mapstructure:"mode" json:"mode"`
	GlobalUrlPrefix string `mapstructure:"global-url-prefix" json:"globalUrlPrefix"`
	Port            int    `mapstructure:"port" json:"port"`
	InitData        bool   `mapstructure:"init-data" json:"initData"`
}

type mysqlConfig struct {
	Username    string `mapstructure:"username" json:"username"`
	Password    string `mapstructure:"password" json:"password"`
	Database    string `mapstructure:"database" json:"database"`
	Host        string `mapstructure:"host" json:"host"`
	Port        int    `mapstructure:"port" json:"port"`
	Query       string `mapstructure:"query" json:"query"`
	LogMode     bool   `mapstructure:"log-mode" json:"logMode"`
	TablePrefix string `mapstructure:"table-prefix" json:"tablePrefix"`
	Charset     string `mapstructure:"charset" json:"charset"`
	Collation   string `mapstructure:"collation" json:"collation"`
}

type logsConfig struct {
	Level      zapcore.Level `mapstructure:"level" json:"level"`            // 日志级别
	Path       string        `mapstructure:"path" json:"path"`              // 日志存放目录
	MaxSize    int           `mapstructure:"max-size" json:"maxSize"`       // 文件最大大小
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"` // 最大备份数
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`         // 存放时间，单位天
	Compress   bool          `mapstructure:"compress" json:"compress"`      // 是否开启压缩
}

type jwtConfig struct {
}

type redisConfig struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
