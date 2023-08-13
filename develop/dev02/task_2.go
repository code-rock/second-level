package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var isLetter = regexp.MustCompile(`[a-zA-Z]`).MatchString
var isSlash = regexp.MustCompile(`\\`).MatchString
var isNumber = regexp.MustCompile(`[0-9]`).MatchString

type sStringUnpackage struct {
	newString    strings.Builder
	slashCounter int
	repidable    string
}

func (su *sStringUnpackage) letterProcessing(char string) {
	su.newString.WriteString(char)
	su.slashCounter = 0
	su.repidable = char
}

func (su *sStringUnpackage) slashProcessing(char string) {
	su.slashCounter += 1
	su.repidable = char
}

func (su *sStringUnpackage) numberProcessing(char string) {
	if n, ok := strconv.ParseInt(char, 10, 8); ok == nil {
		n := int(n)

		if su.slashCounter > 0 {
			if su.slashCounter%2 == 0 {
				su.newString.WriteString(strings.Repeat(`\`, n+su.slashCounter-2))
			} else {

				if _, ok := strconv.ParseInt(string(su.repidable), 10, 8); ok == nil {
					su.newString.WriteString(strings.Repeat(su.repidable, n))
					su.repidable = ""
				} else {
					su.repidable = char
					su.newString.WriteString(char)
				}

			}
		} else {
			if n-1 > 0 {
				su.newString.WriteString(strings.Repeat(su.repidable, n-1))
			} else {
				su.newString.WriteString("")
			}
			su.repidable = ""
		}
		su.slashCounter = 0
	}
}

func (su *sStringUnpackage) unpacking(s string) string {
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if isLetter(char) {
			su.letterProcessing(char)
		} else if isSlash(char) {
			su.slashProcessing(char)
		} else if isNumber(char) {
			su.numberProcessing(char)
		}
	}
	return su.newString.String()
}

func stringFactory() *sStringUnpackage {
	return &sStringUnpackage{}
}

func StringUnpacking(str string) string {
	s := stringFactory()
	return s.unpacking(str)
}

func main() {
	fmt.Println(StringUnpacking(`qwe\4\5`))
}
