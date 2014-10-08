// STARTIMPORT OMIT

package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	out = flag.String("o", "index.json", "path to index file")
)

// ENDIMPORT OMIT

// STARTINTERFACE OMIT

type encoder interface {
	Encode(data interface{}) error
}

func newEncoder(name string, w io.Writer) encoder {
	switch {
	case strings.HasSuffix(name, ".json"):
		return json.NewEncoder(w)
	case strings.HasSuffix(name, ".xml"):
		return xml.NewEncoder(w)
	}
	return json.NewEncoder(w)
}

// ENDINTERFACE OMIT

func main() {
	var err error
	flag.Parse()

	// STARTWALK OMIT
	stats := make([]Stat, 0, 10)
	for _, root := range flag.Args() { // HLxxx
		if root == "" {
			root = "."
		}
		root, err = filepath.Abs(os.ExpandEnv(root))
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}

		s, err := index(root) // HLxxx
		if err != nil {
			log.Fatalf("error: %v\n", err)
		}
		stats = append(stats, s...) // HLxxx
	}
	// ENDWALK OMIT

	// STARTENCODE OMIT
	f, err := os.Create(*out)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
	defer f.Close() // HLxxx

	err = newEncoder(f.Name(), f).Encode(&stats) // HLxxx
	if err != nil {
		log.Fatalf("error encoding stats: %v\n", err)
	}

	err = f.Close() // HLxxx
	if err != nil {
		log.Fatalf("error closing file: %v\n", err)
	}
	// ENDENCODE OMIT
}

// STARTSTAT OMIT

// Stat holds metadata about dirs and files
type Stat struct {
	Name string // full path to object
	Type Type   // type of object (Dir|File)
	Size int64  // size of object in bytes
}

// Type describes the type of a filesystem object (dir|file)
type Type int

// possible types of filesystem objects
const (
	Invalid Type = 0
	Dir     Type = 1
	File    Type = 2
)

// ENDSTAT OMIT

// STARTINDEX OMIT
func index(root string) ([]Stat, error) {
	stats := make([]Stat, 0, 10) // HLxxx
	err := filepath.Walk(root, func(path string, fi os.FileInfo, err error) error {
		fmt.Printf(">>> %s\n", path)
		stat := Stat{
			Name: path,
			Size: fi.Size(),
		}
		switch {
		case fi.IsDir():
			stat.Type = Dir
		case fi.Mode().IsRegular():
			stat.Type = File
		}

		stats = append(stats, stat) // HLxxx

		if fi.IsDir() {
			return nil
		}
		return err
	})

	if err != nil {
		return nil, err
	}

	return stats, err
}

// ENDINDEX OMIT
