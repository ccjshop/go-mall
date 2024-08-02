package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config 配置
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		DB   `yaml:"db"`
		Oss  `yaml:"oss"`
		Jwt  `yaml:"jwt"`
	}

	// App app配置
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP http配置
	HTTP struct {
		IP   string `env-required:"true" yaml:"ip" env:"HTTP_IP"`
		Port uint32 `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log 日志配置
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// DB 数据库配置
	DB struct {
		Username string `env-required:"true" yaml:"username" env:"DB_USERNAME"`
		Password string `env-required:"true" yaml:"password" env:"DB_PASSWORD"`
		Host     string `env-required:"true" yaml:"host" env:"DB_HOST"`
		Port     uint32 `env-required:"true" yaml:"port" env:"DB_PORT"`
		DbName   string `env-required:"true" yaml:"db_name" env:"DB_DB_NAME"`
		Timeout  string `env-required:"true" yaml:"timeout" env:"DB_TIMEOUT"`
	}

	Oss struct {
		BaseUrl string `env-required:"true" yaml:"base_url" env:"OSS_BaseUrl"`
	}

	Jwt struct {
		TimeOut     uint32   `yaml:"time_out"`     // 超时时间，s
		Issuer      string   `yaml:"issuer"`       // 签证签发人
		SignKey     string   `yaml:"sign_key"`     // 密钥
		TokenHeader string   `yaml:"token_header"` // jwt请求头
		TokenHead   string   `yaml:"token_head"`   //
		Whitelist   []string `yaml:"whitelist"`    // 白名单
		Blacklist   []string `yaml:"blacklist"`    // 黑名单
	}
)

// NewConfig 返回程序配置
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
