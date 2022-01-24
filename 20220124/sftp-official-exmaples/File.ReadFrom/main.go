package main

import (
	"bufio"
	"io"

	"github.com/pkg/sftp"
)

func main() {
	// Using Bufio to buffer writes going to an sftp.File won't buffer as it
	// skips buffering if the underlying writer support ReadFrom. The
	// workaround is to wrap your writer in a struct that only implements
	// io.Writer.
	//
	// For background see github.com/pkg/sftp/issues/125

	var data_source io.Reader
	var f *sftp.File
	type writerOnly struct{ io.Writer }
	bw := bufio.NewWriter(writerOnly{f}) // no ReadFrom()
	bw.ReadFrom(data_source)
}
