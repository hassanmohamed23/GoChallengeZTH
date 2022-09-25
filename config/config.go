package config

import (
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"go.uber.org/zap"
)

// DatabaseConfigurations database configurations
type DatabaseConfigurations struct {
	Dsn string `koanf:"dsn"`
}

// Kafka configurations
type KafkaConfiguration struct {
	Brokers   []string `koanf:"broker"`
	Topic     string   `koanf:"topic"`
	MaxBytes  int      `koanf:"maxbytes"`
	Network   string   `koanf:"network"`
	Partition int      `koanf:"partition"`
}

// Configurations Application wide configurations
type Configurations struct {
	Database DatabaseConfigurations `koanf:"database"`
	Kafka    KafkaConfiguration     `koanf:"kafka"`
}

/*func LoadConfig() *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/config.yml"), yaml.Parser())
	if err != nil {
		log.Error.Printf("Failed to locate configurations. %v", err)
	}
	var configuration Configurations
	err = k.Unmarshal("", &configuration)
	if err != nil {
		log.Error.Printf("Failed to load configurations. %v", err)
	}
	return &configuration
}*/
func LoadConfig(logger *zap.SugaredLogger) *Configurations {
	k := koanf.New(".")
	err := k.Load(file.Provider("resources/environment/config.yml"), yaml.Parser())
	if err != nil {
		logger.Fatalf("Failed to locate configurations. %v", err)
	}
	// Searches for env variables and will transform them into koanf format
	// e.g. SERVER_PORT variable will be server.port: value
	err = k.Load(env.Provider("", ".", func(s string) string {
		return strings.Replace(strings.ToLower(s), "_", ".", -1)
	}), nil)
	if err != nil {
		logger.Fatalf("Failed to replace environment variables. %v", err)
	}
	var configuration Configurations
	err = k.Unmarshal("", &configuration)
	if err != nil {
		logger.Fatalf("Failed to load configurations. %v", err)
	}
	return &configuration
}
