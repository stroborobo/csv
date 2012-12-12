package main

import (
	"flag"
	"os"
	"io"
	"fmt"
	"encoding/csv"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

func main() {
	var encoding string
	flag.StringVar(&encoding, "enc", "", "input encoding, e.g. latin9, defaults to UTF-8")

	flag.Parse()

	var f *os.File
	var err error
	if len(flag.Args()) != 0 {
		f, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(10)
		}
	} else {
		f = os.Stdin
	}
	var inputReader io.Reader
	if encoding != "" {
		inputReader, err = charset.NewReader(encoding, f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(20)
		}
	} else {
		inputReader = f
	}
	r := csv.NewReader(inputReader)
	r.Comma = ';'
	r.TrailingComma = true
	r.TrimLeadingSpace = true

	data, err := r.ReadAll()
	if len(os.Args) == 2 {
		f.Close()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(30)
	}
	if len(data) == 0 || len(data[0]) == 0 {
		os.Exit(0)
	}

	colLens := make(map[int]int)
	for _, row := range data {
		for i, col := range row {
			cl := len(col)
			l, ex := colLens[i]
			if !ex || cl > l {
				colLens[i] = cl
			}
		}
	}

	for _, row := range data {
		for i, col := range row {
			fmt.Printf(fmt.Sprint("%-", colLens[i] + 1, "s"), col)
			if i != len(colLens) {
				fmt.Print("| ")
			}
		}
		fmt.Print("\n")
	}
}
