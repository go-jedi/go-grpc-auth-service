package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type LoggerConfig struct {
	IsJSON     bool   `yaml:"is_json"`
	AddSource  bool   `yaml:"add_source"`
	Level      string `yaml:"level"`
	SetFile    bool   `yaml:"set_file"`
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

type JWTConfig struct {
	SecretPath    string `yaml:"secret_path"`
	SecretHashLen int    `yaml:"secret_hash_len"`
	AccessExpAt   int    `yaml:"access_exp_at"`
	RefreshExpAt  int    `yaml:"refresh_exp_at"`
}

type PostgresConfig struct {
	Host          string `yaml:"host"`
	User          string `yaml:"user"`
	Password      string `yaml:"password"`
	DBName        string `yaml:"dbname"`
	Port          int    `yaml:"port"`
	SSLMode       string `yaml:"sslmode"`
	PoolMaxConns  int    `yaml:"pool_max_conns"`
	MigrationsDir string `yaml:"migrations_dir"`
}

type RedisConfig struct {
	Addr            string `yaml:"addr"`
	Password        string `yaml:"password"`
	DB              int    `yaml:"db"`
	DialTimeout     int    `yaml:"dial_timeout"`
	ReadTimeout     int    `yaml:"read_timeout"`
	WriteTimeout    int    `yaml:"write_timeout"`
	PoolSize        int    `yaml:"pool_size"`
	MinIdleConns    int    `yaml:"min_idle_conns"`
	PoolTimeout     int    `yaml:"pool_timeout"`
	MaxRetries      int    `yaml:"max_retries"`
	MinRetryBackoff int    `yaml:"min_retry_backoff"`
	MaxRetryBackoff int    `yaml:"max_retry_backoff"`
}

type GRPCServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Config struct {
	Logger     LoggerConfig     `yaml:"logger"`
	JWT        JWTConfig        `yaml:"jwt"`
	Postgres   PostgresConfig   `yaml:"postgres"`
	Redis      RedisConfig      `yaml:"redis"`
	GRPCServer GRPCServerConfig `yaml:"grpcserver"`
}

// LoadConfig load config file.
func LoadConfig() string {
	var cf string

	flag.StringVar(&cf, "config", "config.yaml", "config file path")
	flag.Parse()

	return cf
}

// ParseConfig parse config file.
func ParseConfig(configFile string) (config Config, err error) {
	f, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&config)

	return config, err
}

// GetConfig get config.
func GetConfig() (config Config, err error) {
	cf := LoadConfig()

	return ParseConfig(cf)
}
