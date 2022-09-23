package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	g := gin.Default()
	g.Handle("GET", "/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	server := NewServer(g, ":8080", func(server *Server) {
		server.ShutdownTimeout = time.Second
	})

	go func() {
		time.Sleep(2 * time.Second)

		resp, err := http.Get("http://localhost:8080")
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		time.Sleep(1 * time.Second)
		err = server.Shutdown()
		assert.NoError(t, err)
		t.Log("shutdown completed")
	}()

	_ = server.Start()
}
