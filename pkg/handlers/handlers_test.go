// handlers/handlers_test.go

package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
	"hotel/middleware"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

// Mocking function for testing
func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.POST("/login", Login)
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware())
    {
        protected.GET("/protected-resource", ProtectedResource)
    }
    return router
}

func TestLogin(t *testing.T) {
    router := SetupRouter()

    payload := `{"username": "testuser", "password": "password123"}`
    req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(payload)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    _, tokenExists := response["token"]
    assert.True(t, tokenExists)
}

func TestProtectedResource(t *testing.T) {
    router := SetupRouter()

    // First, log in to get a token
    payload := `{"username": "testuser", "password": "password123"}`
    req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(payload)))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    var response map[string]string
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    token, tokenExists := response["token"]
    assert.True(t, tokenExists)

    // Use the token to access the protected resource
    req, _ = http.NewRequest("GET", "/protected-resource", nil)
    req.Header.Set("Authorization", "Bearer "+token)
    w = httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    var protectedResponse map[string]interface{}
    err = json.Unmarshal(w.Body.Bytes(), &protectedResponse)
    assert.NoError(t, err)
    message, messageExists := protectedResponse["message"]
    assert.True(t, messageExists)
    assert.Equal(t, "Hello, authenticated user!", message)
}
