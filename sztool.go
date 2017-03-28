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
// The out parameter can be omitted.  For compression it defaults to
// in.sz, where in is the input file.  For decompression it defaults
// to stdout.

package main

import (
	"flag"
	"fmt"
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
	bufsize int = 64 * 1024
)

func errclose(c io.Closer) {
	err := c.Close()
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%v", err))
	}
}

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
		_, err = out.Write(buf[0:n])
		if err != nil {
			panic(err)
		}
	}
}

func compress() {

	wtr := snappy.NewBufferedWriter(out)
	defer errclose(wtr)
	buf := make([]byte, bufsize)

	for {
		n, err := in.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		_, err = wtr.Write(buf[0:n])
		if err != nil {
			panic(err)
		}
	}
}

func main() {

	dec := flag.Bool("d", false, "decompress")
	cmp := flag.Bool("c", false, "compress")
	flag.Parse()

	if *dec && *cmp {
		os.Stderr.WriteString("sztool: can't set both -d and -c\n")
		os.Exit(1)
	}
	if !(*dec || *cmp) {
		os.Stderr.WriteString("sztool: either -d or -c must be set\n")
		os.Exit(1)
	}

	narg := len(flag.Args())
	if narg < 1 || narg > 2 {
		os.Stderr.WriteString("sztool: wrong number of arguments\n")
		os.Exit(1)
	}

	infile := flag.Args()[0]
	if *cmp && strings.HasSuffix(infile, ".sz") {
		os.Stderr.WriteString("sztool: input is already compressed\n")
		os.Exit(1)
	}

	// Get the output file name (or use stdout).
	var outfile string
	if narg == 2 {
		outfile = flag.Args()[1]
	} else {
		if *dec {
			outfile = "-" // use stdout
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
		defer errclose(in)
	}

	// Setup output io.Writer
	if outfile == "-" {
		out = os.Stdout
	} else {
		out, err = os.Create(outfile)
		if err != nil {
			panic(err)
		}
		defer errclose(out)
	}

	if *dec {
		decompress()
	} else if *cmp {
		compress()
	}
}
