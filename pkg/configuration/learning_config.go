package configuration

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Props        *Config
	configLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)
)

type Config struct {
	DB struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"db"`
	RabbitMQ struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"rabbitmq"`
	Gin struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"gin"`
	Jwt struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
	Mongo struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Dbname   string `yaml:"dbname"`
		Port     string `yaml:"port"`
	} `yaml:"mongo"`
	Kafka struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"kafka"`
}

func LoadConfig(filePath string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		configLogger.Fatalf("Error reading properties. Error: %v", err)
	}

	var config Config
	_ = yaml.Unmarshal(file, &config)
	Props = &config
}
