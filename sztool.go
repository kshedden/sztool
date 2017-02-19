// sztool is a simple command-line tool for snappy compressed files.
//
// The basic usage for compression is:
//    sztool -c in out
//
// The basic usage for decompression is:
//    sztool -d in out
//
// Input and output can be specified as "-" to obtain stdin or stdout.
//
// The out parameter can be omitted, for compression it defaults to
// in.sz, for decompression it defaults to stdout.

package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"github.com/golang/snappy"
)

var (
	in  io.ReadCloser
	out io.WriteCloser
)

const (
	bufsize int = 1024
)

func decompress() {

	rdr := snappy.NewReader(in)
	buf := make([]byte, bufsize)

	for {
		n, err := rdr.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		out.Write(buf[0:n])
	}
}

func compress() {

	wtr := snappy.NewBufferedWriter(out)
	defer wtr.Close()
	buf := make([]byte, bufsize)

	for {
		n, err := in.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		wtr.Write(buf[0:n])
	}
}

func main() {

	dec := flag.Bool("d", false, "decompress")
	cmp := flag.Bool("c", false, "compress")
	flag.Parse()

	if *dec && *cmp {
		print("can't set both -d and -c")
		os.Exit(1)
	}

	narg := len(flag.Args())
	if narg != 1 && narg != 2 {
		print("wrong number of arguments")
		os.Exit(1)
	}

	infile := flag.Args()[0]
	if *cmp && strings.HasSuffix(infile, ".sz") {
		print("input is already compressed\n")
		os.Exit(1)
	}
	var outfile string
	if narg == 2 {
		outfile = flag.Args()[1]
	} else {
		if *dec {
			outfile = "-"
		} else {
			outfile = infile + ".sz"
		}
	}

	// Setup input io.Reader
	var err error
	if infile == "-" {
		in = os.Stdin
	} else {
		in, err = os.Open(infile)
		if err != nil {
			panic(err)
		}
		defer in.Close()
	}

	// Setup output io.Writer
	if outfile == "-" {
		out = os.Stdout
	} else {
		out, err = os.Create(outfile)
		if err != nil {
			panic(err)
		}
		defer out.Close()
	}

	if *dec {
		decompress()
	} else if *cmp {
		compress()
	}
}
