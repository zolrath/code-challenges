package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// decodeMorse decodes an individual letter from morse to alphanumeric.
func decodeMorse(m string) rune {
	morse := map[string]rune{
		// Letters
		".-":   'A',
		"-...": 'B',
		"-.-.": 'C',
		"-..":  'D',
		".":    'E',
		"..-.": 'F',
		"--.":  'G',
		"....": 'H',
		"..":   'I',
		".---": 'J',
		"-.-":  'K',
		".-..": 'L',
		"--":   'M',
		"-.":   'N',
		"---":  'O',
		".--.": 'P',
		"--.-": 'Q',
		".-.":  'R',
		"...":  'S',
		"-":    'T',
		"..-":  'U',
		"...-": 'V',
		".--":  'W',
		"-..-": 'X',
		"-.--": 'Y',
		"--..": 'Z',
		// Numbers
		".----": '1',
		"..---": '2',
		"...--": '3',
		"....-": '4',
		".....": '5',
		"-....": '6',
		"--...": '7',
		"---..": '8',
		"----.": '9',
		"-----": '0',
		// Space
		"_": ' ',
	}
	return morse[m]
}

func decodeMorseLine(line string) string {
	line = strings.Replace(line, "  ", " _ ", -1)
	elems := strings.Split(line, " ")
	alpha := make([]rune, len(elems))
	for i, v := range elems {
		alpha[i] = decodeMorse(v)
	}
	return string(alpha)
}

func main() {
	f, _ := os.Open(os.Args[1])
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(decodeMorseLine(s.Text()))
	}
}
