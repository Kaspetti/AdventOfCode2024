package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)


func main() {
    fmt.Printf("Task01: %d\n", task01())
}


func task01() int {
    input, err := os.ReadFile("input")
    if err != nil {
        panic(err)
    }
    inputStr := string(input)

    r, err := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
    if err != nil {
        panic(err)
    }

    total := 0
    matches := r.FindAllString(inputStr, -1)
    for _, match := range matches {
        split := strings.Split(match, ",")
        num0Str := split[0][4:]
        num1Str := split[1][:len(split[1])-1]

        num0, err := strconv.Atoi(num0Str)
        if err != nil {
            panic(err)
        }
        num1, err := strconv.Atoi(num1Str)
        if err != nil {
            panic(err)
        }

        total += num0 * num1
    }

    return total
}
