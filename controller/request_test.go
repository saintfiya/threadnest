package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPageInfoDefaultsAndBounds(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name     string
		query    string
		wantPage int64
		wantSize int64
	}{
		{name: "missing", wantPage: 1, wantSize: 10},
		{name: "invalid", query: "?page=-3&size=abc", wantPage: 1, wantSize: 10},
		{name: "maximum", query: "?page=2&size=1000", wantPage: 2, wantSize: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Request = httptest.NewRequest("GET", "/posts"+tt.query, nil)
			page, size := getPageInfo(ctx)
			if page != tt.wantPage || size != tt.wantSize {
				t.Fatalf("getPageInfo() = (%d, %d), want (%d, %d)", page, size, tt.wantPage, tt.wantSize)
			}
		})
	}
}
