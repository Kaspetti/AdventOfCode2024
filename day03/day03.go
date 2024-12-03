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
    fmt.Printf("Task02: %d\n", task02())
}


func task01() int {
    input, err := os.ReadFile("input")
    if err != nil {
        panic(err)
    }
    inputStr := string(input)

    r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

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


func task02() int {
    input, err := os.ReadFile("input")
    if err != nil {
        panic(err)
    }
    inputStr := string(input)

    r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do(n't)?\(\)`)
    matches := r.FindAllString(inputStr, -1)

    total := 0
    do := true
    for _, match := range matches {
        if match == "do()" {
            do = true
            continue
        } else if match == "don't()" {
            do = false
            continue
        }

        if !do {
            continue
        }

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
