package qspew

import (
	"bytes"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Dump the given value to stderr.
func Dump(v ...any) {
	var buf bytes.Buffer
	spew.Fdump(&buf, v...)
	buf.WriteTo(os.Stderr)
}
