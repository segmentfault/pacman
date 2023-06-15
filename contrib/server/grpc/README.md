# grpc server
> grpc server implemented by [grpc](https://github.com/grpc/grpc)

## Usage
```go
package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestNewServer(t *testing.T) {
	addr := "0.0.0.0:9090"
	grpcServer := grpc.NewServer()
	server := NewServer(grpcServer, addr)

	go func() {
		ctx, cancelFunc := context.WithTimeout(context.TODO(), time.Second*3)
		defer cancelFunc()
		_, err := grpc.DialContext(
			ctx,
			"127.0.0.1:9090",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		assert.NoError(t, err)

		time.Sleep(1 * time.Second)
		err = server.Shutdown()
		assert.NoError(t, err)
		t.Log("shutdown completed")
	}()

	_ = server.Start()
}


```
