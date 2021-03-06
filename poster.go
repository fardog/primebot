package primebot

import (
	"context"
	"fmt"
	"io"
	"os"
)

type Poster interface {
	Post(context.Context, uint64) error
}

func NewWriterPoster(w io.Writer) *WritePoster {
	return &WritePoster{w: w}
}

func NewStdoutPoster() *WritePoster {
	return &WritePoster{w: os.Stdout}
}

type WritePoster struct {
	w io.Writer
}

func (w *WritePoster) Post(ctx context.Context, status uint64) error {
	_, err := fmt.Fprintf(w.w, "%d\n", status)
	return err
}
