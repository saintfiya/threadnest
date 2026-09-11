package mysql

import "testing"

func TestPasswordVerification(t *testing.T) {
	const password = "correct horse battery staple"
	hash, err := hashPassword(password)
	if err != nil {
		t.Fatalf("hashPassword() error = %v", err)
	}

	valid, needsUpgrade := verifyPassword(hash, password)
	if !valid || needsUpgrade {
		t.Fatalf("bcrypt password: valid=%v needsUpgrade=%v", valid, needsUpgrade)
	}
	valid, _ = verifyPassword(hash, "wrong")
	if valid {
		t.Fatal("verifyPassword() accepted an incorrect password")
	}
}

func TestLegacyPasswordRequiresUpgrade(t *testing.T) {
	const password = "legacy-password"
	valid, needsUpgrade := verifyPassword(legacyPasswordHash(password), password)
	if !valid || !needsUpgrade {
		t.Fatalf("legacy password: valid=%v needsUpgrade=%v", valid, needsUpgrade)
	}
}
