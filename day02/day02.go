package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
    fmt.Printf("Task 01: %d\n", task01())
    fmt.Printf("Task 02: %d\n", task02())
}


func task01() int {
    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    
    totalSafe := 0

    lineLoop:
    for scanner.Scan() {
        line := scanner.Text()
        levels := strings.Split(line, " ")
        levelsInt := make([]int, 0)
        for _, level := range levels {
            levelInt, err := strconv.Atoi(level)
            if err != nil {
                panic(err)
            }
            levelsInt = append(levelsInt, levelInt)
        }

        previous := levelsInt[0]
        direction := "up"
        if levelsInt[1] < levelsInt[0] {
            direction = "down"
        }

        for _, level := range levelsInt[1:] {
            // Check direction
            if level > previous && direction == "down" {
                continue lineLoop
            }
            if level < previous && direction == "up" {
                continue lineLoop
            }

            diff := previous - level
            if diff == 0 || absI(diff) > 3 {
                continue lineLoop
            }

            previous = level
        }

        totalSafe += 1
    }
    
    return totalSafe
}


func task02() int {
    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    
    totalSafe := 0

    lineLoop: 
    for scanner.Scan() {
        line := scanner.Text()
        levels := strings.Split(line, " ")
        levelsInt := make([]int, 0)
        for _, level := range levels {
            levelInt, err := strconv.Atoi(level)
            if err != nil {
                panic(err)
            }
            levelsInt = append(levelsInt, levelInt)
        }
        
        validUp, iUp := isValidUp(levelsInt)
        validDown, iDown := isValidDown(levelsInt)

        if validUp || validDown {
            totalSafe += 1
            continue
        }

        // Check removing iUp+-
        for i := iUp-1; i <= iUp+1; i++ {
            if i < 0 || i >= len(levelsInt) {
                continue
            }

            newLevels := make([]int, len(levelsInt))
            copy(newLevels, levelsInt)
            newLevels = append(newLevels[:i], newLevels[i+1:]...)
            if valid, _ := isValidUp(newLevels); valid {
                totalSafe += 1
                continue lineLoop
            }
        }

        // Check removing iDown+-
        for i := iDown-1; i <= iDown+1; i++ {
            if i < 0 || i >= len(levelsInt) {
                continue
            }

            newLevels := make([]int, len(levelsInt))
            copy(newLevels, levelsInt)
            newLevels = append(newLevels[:i], newLevels[i+1:]...)
            if valid, _ := isValidDown(newLevels); valid {
                totalSafe += 1
                continue lineLoop
            }
        }
    }

    return totalSafe
}


func absI(n int) int {
    if n < 0 {
        return -n
    }

    return n
}


func isValidUp(levels []int) (bool, int) {
    for i, level := range levels[1:] {
        if level < levels[i] {
            return false, i+1
        }

        diff := level - levels[i]
        if diff == 0 || absI(diff) > 3 {
            return false, i+1
        }
    }

    return true, -1
}


func isValidDown(levels []int) (bool, int) {
    for i, level := range levels[1:] {
        if level > levels[i] {
            return false, i+1
        }

        diff := level - levels[i]
        if diff == 0 || absI(diff) > 3 {
            return false, i+1
        }
    }

    return true, -1
}
