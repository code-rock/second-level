package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func Cut(field int, delimiter string, separated bool) {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		line := scanner.Text()
		cols := strings.Split(line, delimiter)
		if separated && len(cols) <= 1 {
			continue
		} else {
			if field > 0 && field < len(cols) {
				io.WriteString(os.Stdout, fmt.Sprintf("%s \n", cols[field]))
			} else {
				io.WriteString(os.Stdout, fmt.Sprintf("%v \n", cols))
			}
		}
	}
}

func main() {

	field := flag.Int("f", 0, "Select fields (columns).")
	delimiter := flag.String("d", "	", "Use a different delimiter.")
	separated := flag.Bool("s", false, "Only strings with delimiter.")
	flag.Parse()

	Cut(*field, *delimiter, *separated)
}
