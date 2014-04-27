package main

import (
	"flag"
	"os"
	"io"
	"fmt"
	"strings"
	"unicode/utf8"
	"encoding/csv"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

func main() {
	var fileEncoding,
	    outputEncoding,
	    parseSeperator,
	    printSeperator string
	var debug bool
	flag.StringVar(&fileEncoding, "e", "", "input encoding, e.g. latin9, defaults to UTF-8")
	flag.StringVar(&outputEncoding, "o", "", "output encoding, e.g. latin9, defaults to LC_ALL/LANG or UTF-8")
	flag.StringVar(&parseSeperator, "c", ";", "seperator char used for parsing")
	flag.StringVar(&printSeperator, "s", "|", "seperator string used for printing")
	flag.BoolVar(&debug, "d", false, "debug output")
	// TODO
	//var alignRight bool
	//flag.BoolVar(&alignRight, "r", false, "align values to the right instead to the left")

	flag.Parse()

	if utf8.RuneCountInString(parseSeperator) > 1 {
		fmt.Fprintln(os.Stderr, "The parse seperator must be a single char.")
		flag.Usage()
		os.Exit(5)
	}

	if outputEncoding == "" {
		outputEncoding = getOutputEnc()
	}

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
	if fileEncoding != "" {
		inputReader, err = charset.NewReader(fileEncoding, f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "input encoding: %s\n", err)
			os.Exit(20)
		}
	} else {
		inputReader = f
	}
	r := csv.NewReader(inputReader)
	r.Comma, _ = utf8.DecodeLastRuneInString(parseSeperator)
	r.TrailingComma = true
	r.TrimLeadingSpace = true
	r.LazyQuotes = true

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

	if debug {
		fmt.Fprintf(os.Stderr, "DEBUG columns: %d\n", len(data[0]))
	}

	colLens := make(map[int]int)
	for ri, row := range data {
		for ci, col := range row {
			col = strings.Trim(col, " \t")
			data[ri][ci] = col
			cl := utf8.RuneCountInString(col)
			l, ex := colLens[ci]
			if !ex || cl > l {
				colLens[ci] = cl
			}
		}
	}

	var out io.Writer = os.Stdout;
	if outputEncoding != "UTF-8" {
		out, err = charset.NewWriter(outputEncoding, out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "output encoding: %s\n", err)
			os.Exit(40)
		}
	}

	for _, row := range data {
		for i, col := range row {
			fmt.Fprintf(out, fmt.Sprint("%-", colLens[i] + 1, "s"), col)
			if i != len(colLens) - 1 {
				fmt.Fprintf(out, "%s ", printSeperator)
			}
		}
		fmt.Fprint(out, "\n")
	}
}

func getOutputEnc() string {
	env := os.Getenv("LC_ALL")
	if len(env) == 0 {
		env = os.Getenv("LANG")
		if len(env) == 0 {
			return "UTF-8"
		}
	}
	arr := strings.Split(env, ".")
	if len(arr) != 2 {
		return "UTF-8"
	}
	enc := arr[1]
	// alias
	if strings.ToLower(enc) == "iso8859-15" {
		enc = "iso-8859-15"
	}
	return enc
}
