package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AbsInt (x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {

	// create a variable to hold knob condition - dial
	var dial int = 50
	var countZero int = 0

	filePath := "input.txt"
	
	// read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	cleanedContent := strings.TrimSpace(string(content))

	lines := strings.Split(cleanedContent, "\n")
	// fmt.Println("Total lines:", len(lines))
	// fmt.Println("Dial condition:", dial)


	for _, line := range lines {

		var tempCount int = 0

		direction := line[0]

		var clicks, err = strconv.Atoi(line[1:])
		if err != nil {
			// fmt.Println("Error converting rotation value:", err)
			return
		}

		switch direction {
		case 'L':
			dial, tempCount = turnDial(dial, -clicks)
		case 'R':
			dial, tempCount = turnDial(dial, clicks)
		}
		countZero += tempCount
		// fmt.Println("Dial condition after", line, "is", dial)
		// fmt.Println("Count to zero increased by", tempCount)
		// fmt.Println("The dial is rotated", line, "to point at", dial, "it points at 0", tempCount, "times.")
	}
	fmt.Println("Final dial condition:", dial)
	fmt.Println("Number of times dial hit 0:", countZero)

}


func turnDial (startPos int, clicks int) (newPos int, zeroPasses int) {

	N := 100

    fullTurns := AbsInt(clicks / N)
    remainder := clicks % N

    newPos = (startPos + remainder) % N
    if newPos < 0 {
        newPos += N
    }

    zeroPasses = fullTurns


    if newPos == 0 && clicks != 0 {
        zeroPasses++
    } else if clicks > 0 && newPos < startPos { // right turn logic
        zeroPasses++
    } else if clicks < 0 && startPos != 0 && newPos > startPos { //left turn logic
        zeroPasses++
    }

    return newPos, zeroPasses

}




func mainCircular() {

	r := ring.New(100)
	countZero := 0

	for i := 0; i < 100; i++ {
		r.Value = i
		r = r.Next()
	}

	// initial r position
	r = r.Move(50) // point at 50
	fmt.Println("Initial dial condition:", r.Value)

	filePath := "input.txt"
	// read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	cleanedContent := strings.TrimSpace(string(content))

	lines := strings.Split(cleanedContent, "\n")
	fmt.Println("Total lines:", len(lines))

	for _, line := range lines {

		// tempCount := 0

		direction := line[0]

		clicks, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println("Error converting rotation value:", err)
			return
		}

		switch direction {
		case 'L':
			for range clicks {
				r = r.Prev()
				if r.Value == 0 {
					countZero++
				}
			}


		case 'R':
			for range clicks {
				r = r.Next()
				if r.Value == 0 {
					countZero++
				}
			}

		}

	}
	fmt.Println("Final dial condition:", r.Value)
	fmt.Println("Number of times dial hit 0:", countZero)


}
