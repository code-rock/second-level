package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var IsLetter = regexp.MustCompile(`[a-zA-Z]`).MatchString
var IsSlash = regexp.MustCompile(`\\`).MatchString
var IsNumber = regexp.MustCompile(`[0-9]`).MatchString

type SStringUnpackage struct {
	newString    strings.Builder
	slashCounter int
	repidable    string
}

func (su *SStringUnpackage) letterProcessing(char string) {
	su.newString.WriteString(char)
	su.slashCounter = 0
	su.repidable = char
}

func (su *SStringUnpackage) slashProcessing(char string) {
	su.slashCounter += 1
	su.repidable = char
}

func (su *SStringUnpackage) numberProcessing(char string) {
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
			su.newString.WriteString(strings.Repeat(su.repidable, n-1))
		}
		su.slashCounter = 0
	}
}

func (su *SStringUnpackage) unpacking(s string) string {
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if IsLetter(char) {
			su.letterProcessing(char)
		} else if IsSlash(char) {
			su.slashProcessing(char)
		} else if IsNumber(char) {
			su.numberProcessing(char)
		}
	}
	return su.newString.String()
}

func main() {
	s := SStringUnpackage{}
	fmt.Println(s.unpacking(`qwe\4\5`))
}
