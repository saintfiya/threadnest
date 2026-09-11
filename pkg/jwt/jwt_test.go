package jwt

import "testing"

func TestInitRejectsShortSecret(t *testing.T) {
	customSecret = nil
	if err := Init("too-short"); err == nil {
		t.Fatal("Init() accepted a secret shorter than 32 bytes")
	}
}

func TestTokenRoundTrip(t *testing.T) {
	customSecret = nil
	if err := Init("0123456789abcdef0123456789abcdef"); err != nil {
		t.Fatalf("Init() error = %v", err)
	}

	token, err := GenToken(42, "alice")
	if err != nil {
		t.Fatalf("GenToken() error = %v", err)
	}
	claims, err := ParseToken(token)
	if err != nil {
		t.Fatalf("ParseToken() error = %v", err)
	}
	if claims.UserID != 42 || claims.Username != "alice" {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}
