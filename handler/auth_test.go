package handler

import (
	"context"
	"testing"

	"example.com/tech-challange-auth-issuer/models"
)

func TestAuthHandler_EmptyDocument(t *testing.T) {
	input := models.Input{Document: ""}
	_, err := AuthHandler(context.Background(), input)

	if err == nil {
		t.Error("Expected error for empty document")
	}
	if err.Error() != "O campo 'document' é obrigatório" {
		t.Errorf("Expected validation error, got: %v", err)
	}
}

func TestAuthHandler_WhitespaceDocument(t *testing.T) {
	input := models.Input{Document: "   "}
	_, err := AuthHandler(context.Background(), input)

	if err == nil {
		t.Error("Expected error for whitespace document")
	}
}
