package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func openFile(url string) (*os.File, error) {
	file, err := os.Open(url)

	if err != nil {
		fmt.Printf("Error when opening file: %s\n", err)
		return nil, err
	}

	return file, nil
}

func grep(file *os.File, pattern string, after int, before int, ignoreCase bool, lineNum bool, invert bool) {
	fileScanner := bufio.NewScanner(file)
	addition := ""
	if ignoreCase {
		addition += "(?i)"
	}
	substr := regexp.MustCompile(addition + pattern)
	fmt.Println("gf")
	cAfter := after
	cBefore := before
	for i := 1; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		cBefore = i - before
		if cBefore < 0 {
			cBefore = 0
		}
		if (invert && !substr.MatchString(line)) || (!invert && substr.MatchString(line)) {

			fileScanner2 := bufio.NewScanner(file)
			for j := 0; before >= 0 && fileScanner2.Scan(); {
				if j >= cBefore {
					if lineNum {
						io.WriteString(os.Stdout, fmt.Sprintf("%d %s \n", i, fileScanner2.Text()))
					} else {
						io.WriteString(os.Stdout, fileScanner2.Text()+"\n")
					}
				}

			}
			if lineNum {
				io.WriteString(os.Stdout, fmt.Sprintf("%d %s \n", i, line))
			} else {
				io.WriteString(os.Stdout, line+"\n")
			}

			cAfter = after
		} else {
			if after > 0 {
				io.WriteString(os.Stdout, line+"\n")
				cAfter -= 1
			}
		}
	}
}

func main() {
	// var err error

	// // Create an ls command.
	// ls := exec.Command("ls", "-la")

	// // Create a grep command that searches for anything
	// // that contains .go in it's filename.
	// grep := exec.Command("grep", "\\.go")

	// // Set grep's stdin to the output of the ls command.
	// grep.Stdin, err = ls.StdoutPipe()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // Set grep's stdout to os.Stdout
	// grep.Stdout = os.Stdout

	// // Start the grep command first. (The order will be last command first)
	// must(grep.Start())

	// // Run the ls command. (Run calls start and also calls wait)
	// must(ls.Run())

	// // Wait for the grep command to finish.
	// must(grep.Wait())
	after := flag.Int("A", 0, "Print +N lines after match.")
	before := flag.Int("B", 0, "Print +N lines until match.")
	// context := flag.Int("C", 0, "(A+B) print ±N lines around the match.")
	// count := flag.Int("c", 0, "Number of lines.")
	ignoreCase := flag.Bool("i", false, "Ignore case.")
	invert := flag.Bool("v", false, "Instead of match, exclude.")
	// fixed := flag.Bool("F", false, "Exact string match, not a pattern.")
	lineNum := flag.Bool("n", false, "Print line number.")

	flag.Parse()
	pattern := flag.Arg(0)
	filename := flag.Arg(1)
	fmt.Println(pattern, filename)
	if filename != "" && pattern != "" {
		file, err := openFile(filename)
		if err == nil {
			grep(file, pattern, *after, *before, *ignoreCase, *lineNum, *invert)
		} else {
			fmt.Println("Файл не парсится")
		}
	}
}
