package main

import (
	"flag"
	"fmt"
	"hw06/copy"
	"os"
)

var f, t string
var l, o int

func init() {
	flag.StringVar(&f, "from", "", "File to read from")
	flag.StringVar(&t, "to", "", "File to read to")
	flag.IntVar(&l, "limit", 0, "Limit for copy")
	flag.IntVar(&o, "offset", 0, "Offset for copy")
	flag.Parse()

	if f == "" {
		fmt.Println("not set from file")
		os.Exit(0)
	}
	if t == "" {
		fmt.Println("not set to file")
		os.Exit(0)
	}
}

func main() {
	err := copy.Copy(f, t, l, o)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
}
