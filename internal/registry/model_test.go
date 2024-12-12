package registry

import (
	"testing"
	"time"
)

func TestServiceStruct(t *testing.T) {
	now := time.Now()
	service := Service{
		ID:          "123456",
		Name:        "TestService",
		Address:     "127.0.0.1",
		Port:        3005,
		RegistredAt: now,
	}

	if service.ID != "123456" {
		t.Errorf("expected ID to be '123456', got %s", service.ID)
	}

	if service.Name != "TestService" {
		t.Errorf("expected Name to be 'TestService', got %s", service.Name)
	}

	if service.Address != "127.0.0.1" {
		t.Errorf("expected Address to be '127.0.0.1', got %s", service.Address)
	}

	if service.Port != 3005 {
		t.Errorf("expected Port to be 3005, got %d", service.Port)
	}

	if !service.RegistredAt.Equal(now) {
		t.Errorf("expected RegistredAt to be '%v', got %v", now, service.RegistredAt)
	}
}
