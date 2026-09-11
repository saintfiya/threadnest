package logger

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDumpRequestWithoutSecrets(t *testing.T) {
	req := httptest.NewRequest("GET", "/resource?token=query-value", nil)
	req.Header.Set("Authorization", "Bearer top-secret-token")
	req.Header.Set("Cookie", "session=top-secret-cookie")

	dump := string(dumpRequestWithoutSecrets(req))
	if strings.Contains(dump, "top-secret-token") || strings.Contains(dump, "top-secret-cookie") || strings.Contains(dump, "query-value") {
		t.Fatalf("request dump contains a secret: %s", dump)
	}
	if !strings.Contains(dump, "[REDACTED]") {
		t.Fatalf("request dump does not contain redaction marker: %s", dump)
	}
}
