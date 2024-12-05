package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type order string

const (
    BEFORE  order = "BEFORE"
    AFTER   order = "AFTER"
)


func main() {
    fmt.Printf("Task01: %d\n", task01())
    fmt.Printf("Task02: %d\n", task02())
}


func task01() int {
    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    rules := make(map[string]bool)

    readAllRules := false
    total := 0
    for scanner.Scan() {
        if scanner.Text() == "" {
            readAllRules = true
            continue
        }

        if !readAllRules {
            rules[scanner.Text()] = true
            continue
        }

        update := strings.Split(scanner.Text(), ",")
        ordered := true
        outer:
        for i, num0 := range update {
            for _, num1 := range update[i+1:] {
                key := num0 + "|" + num1
                if _, ok := rules[key]; ok {
                    continue
                }

                key = num1 + "|" + num0
                if _, ok := rules[key]; ok {
                    ordered = false
                    break outer
                }
            }
        }

        if ordered {
            mid := update[len(update)/2]
            midInt, err := strconv.Atoi(mid) 
            if err != nil {
                panic(err)
            }

            total += midInt
        }
    }

    return total
}


func task02() int {
    return 0
}
