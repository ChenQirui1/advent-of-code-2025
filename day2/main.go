package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	counter := 0

	//file operations
	filePath := "input.txt"
	// read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	cleanedContent := strings.TrimSpace(string(content))

	lines := strings.Split(cleanedContent, ",")

	// testRange := "11-22"

	for _, line := range lines {
		counter += checkRange(line)
	}

	fmt.Println(counter)

}

func checkRange(rangeString string) int {

	tally := 0

	// testRange := "11-22"

	splitedStrings := strings.Split(rangeString, "-")

	startRange, err := strconv.Atoi(splitedStrings[0])

	if err != nil {
		fmt.Println(err)
	}

	endRange, err := strconv.Atoi(splitedStrings[1])

	if err != nil {
		fmt.Println(err)
	}

	checkID := func(id string) bool {

		// for part 1 where its only two
		// return id[0:len(id)/2] == id[len(id)/2:]

		//more general pattern using regex but pattern does not work, as it containes backref which is not allowin RE2
		// pattern := `^(.*)\1+$`
		// re := regexp.MustCompile(pattern) // compile regex
		// return re.MatchString(id)

		// using string repeat

		stringLen := len(id)

		for i := 1; i <= stringLen/2; i++ {
			//check for even
			if stringLen%i == 0 {
				testSubString := id[0:i]
				// fmt.Println(testSubString)
				if strings.Repeat(testSubString, stringLen/i) == id {
					return true
				}
			}
		}
		return false
	}

	for i := startRange; i <= endRange; i++ {

		id := strconv.Itoa(i)

		if checkID(id) {

			tally += i
		}
		// fmt.Println(checkID(id))
	}

	fmt.Println("count: ", tally)

	return tally
}

func checkID(id string) bool {

	// for part 1 where its only two
	// return id[0:len(id)/2] == id[len(id)/2:]

	//more general pattern using regex but pattern does not work, as it containes backref which is not allowin RE2
	// pattern := `^(.*)\1+$`
	// re := regexp.MustCompile(pattern) // compile regex
	// return re.MatchString(id)

	// using string repeat
	stringLen := len(id)

	for i := 1; i <= stringLen/2; i++ {
		//check for even
		if stringLen%i == 0 {
			testSubString := id[0:i]
			fmt.Println(testSubString)
			if strings.Repeat(testSubString, stringLen/i) == id {
				return true
			}
		}
	}
	return false
}
