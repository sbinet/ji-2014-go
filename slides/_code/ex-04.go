package main

import (
	"io"
	"os"

	"github.com/ji-go/rot13" // HLxxx
)

func main() {
	r := rot13.NewReader(os.Stdin)
	io.Copy(os.Stdout, r)
}
