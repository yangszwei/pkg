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

func TestWriter(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	const testStr = "Hello, World!"

	t.Run("should pass", func(t *testing.T) {
		buf := new(bytes.Buffer)
		w := ctxio.NewWriter(ctx, buf)
		n, err := w.Write([]byte(testStr))
		require.NoError(t, err)
		assert.Equal(t, len(testStr), n)
		assert.Equal(t, testStr, buf.String())
	})

	t.Run("should fail", func(t *testing.T) {
		time.Sleep(40 * time.Millisecond)
		buf := new(bytes.Buffer)
		w := ctxio.NewWriter(ctx, buf)
		n, err := w.Write([]byte(testStr))
		assert.ErrorIs(t, err, context.DeadlineExceeded)
		assert.Zero(t, n)
		assert.Zero(t, buf.String())
	})
}
