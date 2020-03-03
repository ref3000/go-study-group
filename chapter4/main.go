package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// go-cutコマンドを実装しよう
func main() {
	var (
		d = flag.String("d", "\t", "delimiter")
		f = flag.Int("f", 0, "fields")
	)
	flag.Parse()

	if *f == 0 {
		panic("-f: fields parameter is required.")
	}

	args := flag.Args()

	if len(args) < 1 {
		panic("arg0: file path is required.")
	}

	path := args[0]
	fp, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	s := bufio.NewScanner(fp)
	for s.Scan() {
		line := s.Text()
		tokens := strings.Split(line, *d)
		var token string
		if len(tokens) >= *f {
			token = tokens[*f-1]
		} else {
			token = ""
		}
		fmt.Println(token)
	}
}
