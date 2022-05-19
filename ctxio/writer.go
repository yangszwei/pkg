package ctxio

import (
	"context"
	"io"
)

// NewWriter creates a new io.Writer that respects the context.
func NewWriter(ctx context.Context, w io.Writer) io.Writer {
	return &writer{ctx: ctx, w: w}
}

// writer wraps an io.Writer to make it respect the context.
type writer struct {
	ctx context.Context
	w   io.Writer
}

// Write writes to the wrapped io.Writer while the context is not done.
func (writer *writer) Write(p []byte) (n int, err error) {
	select {
	case <-writer.ctx.Done():
		return 0, writer.ctx.Err()
	default:
		return writer.w.Write(p)
	}
}
