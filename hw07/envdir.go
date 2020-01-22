package main

import (
	"fmt"
	"hw07/env"
	"os"
)

func init() {

}

func main() {
	if len(os.Args) < 3 {
		_, _ = fmt.Fprintf(os.Stderr, "Not enoth arguments,\n example: envdir ./evndirtest command arg1 arg2\n")
		os.Exit(2)
	}

	environments, err := env.ReadDir(os.Args[1] + "/")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	runStatus := env.RunCmd(os.Args[2:], environments)

	os.Exit(runStatus)
}
