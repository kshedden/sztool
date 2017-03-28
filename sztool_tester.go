package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func TestWriteRead1() {

	fid, err := os.Create("tmp.txt")
	if err != nil {
		panic(err)
	}
	for k := 0; k < 10; k++ {
		c := fmt.Sprintf("%d", k)
		_, err := fid.Write([]byte(strings.Repeat(c, k+1)))
		if err != nil {
			panic(err)
		}
		_, err = fid.Write([]byte("\n"))
		if err != nil {
			panic(err)
		}
	}
	fid.Close()

	cmd := exec.Command("sztool", "-c", "tmp.txt")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("sztool", "-d", "tmp.txt.sz", "tmp2.txt")
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fid, err = os.Open("tmp2.txt")
	if err != nil {
		panic(err)
	}
	rdr := bufio.NewReader(fid)
	for k := 0; k < 10; k++ {
		line, err := rdr.ReadString('\n')
		if err != nil {
			panic(err)
		}
		line = strings.TrimRight(line, "\n")
		c := fmt.Sprintf("%d", k)
		line2 := strings.Repeat(c, k+1)
		if line != line2 {
			_, err := os.Stderr.WriteString(fmt.Sprintf("%s != %s\n", line, line2))
			if err != nil {
				panic(err)
			}
			panic("")
		}
	}
}

func main() {

	TestWriteRead1()
}
