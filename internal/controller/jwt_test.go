package controller

import (
	"log"
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	tests := []struct {
		name  string
		email string
	}{
		{
			name:  "valid",
			email: "test@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateJWT(tt.email)
			if err != nil {
				t.Error("couldn't get token:", err)
			}
			log.Println("token:", token)
			if token == "" {
				t.Error("expected token to not be empty")
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	email := "test@example.com"
	token, err := GenerateJWT(email)
	if err != nil {
		t.Error("couldn't get token:", err)
	}

	tests := []struct {
		name  string
		email string
		exp   time.Duration
	}{
		{
			name:  "valid",
			email: email,
			exp:   time.Hour * 24 * 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims, err := ValidateJWT(token)
			if err != nil {
				t.Error("couldn't validate token:", err)
			}
			if claims.Email != tt.email {
				t.Errorf("expected email to be %s, got %s", tt.email, claims.Email)
			}
			if time.Until(claims.ExpiresAt.Time) > tt.exp {
				t.Errorf("expected token to expire in %s, got %s", tt.exp, time.Until(claims.ExpiresAt.Time))
			}
		})
	}
}
