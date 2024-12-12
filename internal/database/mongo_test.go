package database

import (
	"context"
	"testing"
	"time"
)

func TestConnectMongo_Success(t *testing.T) {
	mockURI := "mongodb://localhost:27017"

	client, err := ConnectMongo(mockURI)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if client == nil {
		t.Error("expected client to be non-nil")
	}

	defer client.Disconnect(context.Background())
}

func TestConnectMongo_InvalidURI(t *testing.T) {
	invalidURI := "invalid-uri"

	client, err := ConnectMongo(invalidURI)
	if err == nil {
		t.Errorf("expected error, got nil")
	}

	if client != nil {
		t.Error("expected client to be nil")
	}
}

func TestConnectMongo_Timeout(t *testing.T) {
	timeoutURI := "mongodb://10.255.255.1:27017"

	start := time.Now()
	_, err := ConnectMongo(timeoutURI)
	elapsed := time.Since(start)

	if err == nil {
		t.Error("expected timeout error, got nil")
	}

	if elapsed < 10*time.Second {
		t.Errorf("expected at least 10 second timeout, got %v", elapsed)
	}
}
