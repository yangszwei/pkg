package ctxio

import (
	"context"
	"io"
)

// NewReader creates a new io.Reader that respects the context.
func NewReader(ctx context.Context, r io.Reader) io.Reader {
	return &reader{ctx: ctx, r: r}
}

// reader wraps an io.Reader to make it respect the context.
type reader struct {
	ctx context.Context
	r   io.Reader
}

// Read reads from the wrapped io.Reader while the context is not done.
func (reader *reader) Read(p []byte) (n int, err error) {
	select {
	case <-reader.ctx.Done():
		return 0, reader.ctx.Err()
	default:
		return reader.r.Read(p)
	}
}
