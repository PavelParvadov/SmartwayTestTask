package config

import (
	"flag"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env" env:"ENV"`
	App App    `yaml:"app"`
	DB  DB     `yaml:"database"`
}

type App struct {
	Host string `yaml:"host" env:"HOST"`
	Port int    `yaml:"port" env:"PORT"`
}

type DB struct {
	Host     string `yaml:"host" env:"DATABASE_HOST"`
	Port     int    `yaml:"port" env:"DATABASE_PORT"`
	Username string `yaml:"username" env:"DATABASE_USERNAME"`
	Password string `yaml:"password" env:"DATABASE_PASSWORD"`
	Name     string `yaml:"db_name" env:"DATABASE_NAME"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		path := fetchConfigPath()
		instance = LoadConfigByPath(path)
	})
	return instance
}

func LoadConfigByPath(path string) *Config {
	var cfg Config

	if path != "" {
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			panic(err)
		}
	}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "config file path")
	if !flag.Parsed() {
		flag.Parse()
	}
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
