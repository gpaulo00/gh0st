package common

import (
	"context"
	"io"
)

// CtxReader is a cancellable reader
type CtxReader struct {
	ctx    context.Context
	reader io.Reader
}

func (r CtxReader) Read(p []byte) (n int, err error) {
	select {
	case <-r.ctx.Done():
		return 0, r.ctx.Err()
	default:
		return r.reader.Read(p)
	}
}

// CtxWriter is a cancellable writer
type CtxWriter struct {
	ctx    context.Context
	writer io.Writer
}

func (w CtxWriter) Write(p []byte) (n int, err error) {
	select {
	case <-w.ctx.Done():
		return 0, w.ctx.Err()
	default:
		return w.writer.Write(p)
	}
}

// NewReader creates a new CtxReader
func NewReader(ctx context.Context, r io.Reader) CtxReader {
	return CtxReader{ctx: ctx, reader: r}
}

// NewWriter creates a new CtxWriter
func NewWriter(ctx context.Context, w io.Writer) CtxWriter {
	return CtxWriter{ctx: ctx, writer: w}
}
