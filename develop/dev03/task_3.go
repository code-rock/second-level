package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mpvl/unique"
)

func openFile(url string) (*os.File, error) {
	file, err := os.Open(url)

	if err != nil {
		fmt.Printf("Error when opening file: %s\n", err)
		return nil, err
	}

	return file, nil
}

func writeFile(lines []string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	for _, v := range lines {
		file.WriteString(v + "\n")
	}

	fmt.Println("qqqq")
}

type SSort struct {
	sortMatrix       [][]string
	sortСolumn       int
	sortByNumber     bool
	sortByMonthName  string
	isUnique         bool
	isReverse        bool
	isIgTaileSpaces  bool
	isSorted         bool
	isSortByNAndSuff bool
	bySorted         bool
}

func (s *SSort) splitToMatrix(f io.Reader) {
	fileScanner := bufio.NewScanner(f)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		s.sortMatrix = append(s.sortMatrix, strings.Fields(line))
	}
}

func (s SSort) joinToLines() []string {
	var lines []string
	for _, v := range s.sortMatrix {
		lines = append(lines, strings.Join(v, " "))
	}
	return lines
}

func (s SSort) removeDuplicateLines(lines []string) []string {
	if s.isUnique {
		unique.Strings(&lines)
	}
	return lines
}

func (s SSort) parseAsN(n1 []string, n2 []string) (int, int, bool) {
	nNew1, err1 := strconv.Atoi(n1[s.sortСolumn])
	nNew2, err2 := strconv.Atoi(n2[s.sortСolumn])
	if err1 == nil && err2 == nil {
		s.bySorted = true
		return nNew1, nNew2, true
	}
	return 0, 0, false
}

func (s SSort) sortedNotification() {
	if s.isSorted {
		if s.bySorted {
			fmt.Println("Sorted")
		} else {
			fmt.Println("Unsorted")
		}
	}
}

func (s SSort) sortReverceN(n1, n2 int) bool {
	if s.isReverse {
		return n1 >= n2
	} else {
		return n1 <= n2
	}
}

func (s SSort) sortReverceStr(n1, n2 string) bool {
	if s.isReverse {
		return n1 >= n2
	} else {
		return n1 <= n2
	}
}

func (s *SSort) sortedMatrix() {
	sort.Slice(s.sortMatrix, func(i, j int) bool {
		if s.sortByNumber {
			n1, n2, ok := s.parseAsN(s.sortMatrix[i], s.sortMatrix[j])
			if ok {
				return s.sortReverceN(n1, n2)
			}
		}
		return s.sortReverceStr(s.sortMatrix[i][s.sortСolumn], s.sortMatrix[j][s.sortСolumn])
	})
}

func (s SSort) sort(url string) {
	f, _ := openFile(url)
	s.splitToMatrix(f)
	s.sortedMatrix()
	lines := s.joinToLines()
	lines = s.removeDuplicateLines(lines)
	writeFile(lines, strings.Replace(url, ".", "_sorted.", 1))
	s.sortedNotification()
}

func main() {

	sortСolumn := flag.Int("k", 0, "Sort by column.")
	sortByNumber := flag.Bool("n", false, "Sort by numeric value.")
	sortByMonthName := flag.String("M", "Aug", "Sort by month name.")
	isUnique := flag.Bool("u", false, "Don't output duplicate lines.")
	isReverse := flag.Bool("r", false, "Sort in reverse order.")
	isIgTaileSpaces := flag.Bool("b", false, "Ignore trailing spaces.")
	isSorted := flag.Bool("c", false, "Check if data is sorted.")
	isSortByNAndSuff := flag.Bool("h", false, "Sort by numerical value considering suffixes.")

	flag.Parse()

	s := SSort{
		sortСolumn:       *sortСolumn,
		sortByNumber:     *sortByNumber,
		sortByMonthName:  *sortByMonthName,
		isUnique:         *isUnique,
		isReverse:        *isReverse,
		isIgTaileSpaces:  *isIgTaileSpaces,
		isSorted:         *isSorted,
		isSortByNAndSuff: *isSortByNAndSuff,
	}

	if filename := flag.Arg(0); filename != "" {
		s.sort(filename)
	}

}
