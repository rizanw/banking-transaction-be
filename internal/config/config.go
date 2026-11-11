package config

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func New(appName string) (*Config, error) {
	fileConfig := getConfigFile(appName)

	f, err := os.Open(fileConfig)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err = yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	// environment variables override the config file for secrets
	if v := os.Getenv("JWT_SECRET"); v != "" {
		cfg.JWT.Secret = v
	}
	if v := os.Getenv("EMAIL"); v != "" {
		cfg.SMTP.Username = v
	}
	if v := os.Getenv("EMAIL_PASSWORD"); v != "" {
		cfg.SMTP.Password = v
	}

	// a JWT secret is required to sign tokens; without one, generate a random
	// key so tokens can't be forged (sessions won't survive a restart).
	if cfg.JWT.Secret == "" {
		secret := make([]byte, 32)
		if _, err = rand.Read(secret); err != nil {
			return nil, fmt.Errorf("generate jwt secret: %w", err)
		}
		cfg.JWT.Secret = string(secret)
		log.Println("[config] JWT_SECRET is not set: using a random key, so tokens are invalidated on restart")
	}
	if len(cfg.JWT.Secret) < 32 {
		log.Println("[config] JWT_SECRET is shorter than 32 bytes: use a longer random value in production")
	}

	if cfg.Redis.Host == "" {
		cfg.Redis.Host = "localhost"
	}
	if cfg.Redis.Port == 0 {
		cfg.Redis.Port = 6379
	}
	if len(cfg.Server.AllowedOrigins) == 0 {
		cfg.Server.AllowedOrigins = []string{"*"}
	}

	return &cfg, nil
}

func getConfigFile(appName string) string {
	var (
		filename = fmt.Sprintf("%s/config.yaml", appName)
	)

	dir, _ := os.Getwd()

	return filepath.Join(dir, "files/etc", filename)
}
