package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(s string) string {
	data, err := os.ReadFile(s)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	contents := strings.Split(readFile("sample-day1.txt"), "\n")
	sum := 0

	for i := 0; i < len(contents); i++ {
		lineChars := strings.Split(contents[i], "")

		var combo []string

		for j := 0; j < len(lineChars); j++ {

			_, err := strconv.Atoi(lineChars[j])

			if err != nil {
				continue
			}

			combo = append(combo, lineChars[j])

		}

		if len(combo) == 1 {

			num, _ := strconv.Atoi(combo[0] + combo[0])

			sum = sum + num
		}

		if len(combo) > 1 {

			num, _ := strconv.Atoi(combo[0] + combo[len(combo)-1])

			sum = sum + num
		}

	}
	println(sum)

}
