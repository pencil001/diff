package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pencil001/diff"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		panic("need two files to diff")
	}

	fileA, fileB := args[0], args[1]

	contentA, err := ioutil.ReadFile(fileA)
	if err != nil {
		panic(err)
	}

	contentB, err := ioutil.ReadFile(fileB)
	if err != nil {
		panic(err)
	}

	linesA := strings.Split(string(contentA), "\n")
	linesB := strings.Split(string(contentB), "\n")

	d := diff.NewMyers(linesA, linesB)

	diffs, err := d.Diff()
	if err != nil {
		panic(err)
	}

	for _, df := range diffs {
		fmt.Println(df)
	}
}
