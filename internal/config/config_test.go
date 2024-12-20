package config

import (
	"os"
	"testing"
)

func TestLoadConfig_WithEnvVariables(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://test-mongo:27017")
	os.Setenv("SERVER_PORT", "3000")
	os.Setenv("DATABASE_NAME", "service_registry")
	os.Setenv("COLLECTION_NAME", "services")

	defer func() {
		os.Unsetenv("MONGO_URI")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("DATABASE_NAME")
		os.Unsetenv("COLLECTION_NAME")
	}()

	cfg := LoadConfig()

	if cfg.MongoURI != "mongodb://test-mongo:27017" {
		t.Errorf("expected MongoURI to be 'mongodb://test-mongo:27017', got '%s'", cfg.MongoURI)
	}

	if cfg.ServerPort != "3000" {
		t.Errorf("expected ServerPort to be '3000', got '%s'", cfg.ServerPort)
	}

	if cfg.DatabaseName != "service_registry" {
		t.Errorf("expected DatabaseName to be 'service_registry', got '%s'", cfg.DatabaseName)
	}

	if cfg.CollectionName != "services" {
		t.Errorf("expected CollectionName to be 'services', got '%s'", cfg.CollectionName)
	}
}

func TestLoadConfig_WithoutEnvVariables(t *testing.T) {
	os.Unsetenv("MONGO_URI")
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("DATABASE_NAME")
	os.Unsetenv("COLLECTION_NAME")

	cfg := LoadConfig()

	if cfg.MongoURI != "mongodb://localhost:27017" {
		t.Errorf("expected MongoURI to be 'mongodb://localhost:27017', got '%s'", cfg.MongoURI)
	}

	if cfg.ServerPort != "3000" {
		t.Errorf("expected ServerPort to be '3000', got '%s'", cfg.ServerPort)
	}

	if cfg.DatabaseName != "service_registry" {
		t.Errorf("expected DatabaseName to be 'service_registry', got '%s'", cfg.DatabaseName)
	}

	if cfg.CollectionName != "services" {
		t.Errorf("expected CollectionName to be 'services', got '%s'", cfg.CollectionName)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "test-value")
	defer os.Unsetenv("TEST_KEY")

	value := getEnv("TEST_KEY", "default-value")
	if value != "test-value" {
		t.Errorf("expected 'test-value', got '%s'", value)
	}

	// Test without environment variable set
	value = getEnv("NON_EXISTENT_KEY", "default-value")
	if value != "default-value" {
		t.Errorf("expected 'default-value', got '%s'", value)
	}
}
