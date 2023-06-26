package types

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresURI string
	OpenAIKey   string
}

func getEnv(e string) string {
	env, exists := os.LookupEnv(e)
	if !exists {
		log.Fatalf("Environment variable %s not found", env)
	}

	return env
}

func getOptionalEnv(e string) string {
	env, _ := os.LookupEnv(e)

	return env
}

func NewConfig(postgresURI, openAIKey string) *Config {
	return &Config{
		PostgresURI: postgresURI,
		OpenAIKey:   openAIKey,
	}
}

func InitConfig() *Config {
	// Try to load .env file in the current working directory.
	// Repeat for parent directories until succeeding, or reaching root.
	cwd, _ := os.Getwd()
	log.Printf("Loading environment in %s", cwd)
	for cwd != "." && cwd != "/" {
		p := filepath.Join(cwd, ".env")
		err := godotenv.Load(p)
		cwd = filepath.Dir(cwd)
		if err == nil {
			log.Printf("Loaded env file: %s", p)
			break
		}
	}

	cfg := NewConfig(
		getEnv("MUSE_POSTGRES_URI"),
		getEnv("MUSE_OPENAI_API_KEY"),
	)

	return cfg
}
