package recovery

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())
	engine.GET("/panic", func(c *gin.Context) {
		panic("test panic")
	})

	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/panic", nil)
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		engine.ServeHTTP(recorder, req)
	})
}
