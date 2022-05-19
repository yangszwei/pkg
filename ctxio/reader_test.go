package ctxio_test

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yangszwei/pkg/ctxio"
	"testing"
	"time"
)

func TestReader(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	const testStr = "Hello, World!"

	t.Run("should pass", func(t *testing.T) {
		r := ctxio.NewReader(ctx, bytes.NewBufferString(testStr))
		buf := make([]byte, len(testStr))
		n, err := r.Read(buf)
		require.NoError(t, err)
		assert.Equal(t, testStr[:n], string(buf))
	})

	t.Run("should fail", func(t *testing.T) {
		time.Sleep(40 * time.Millisecond)
		r := ctxio.NewReader(ctx, bytes.NewBufferString(testStr))
		buf := make([]byte, len(testStr))
		n, err := r.Read(buf)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
		assert.Zero(t, n)
		assert.Equal(t, make([]byte, len(testStr)), buf)
	})
}
