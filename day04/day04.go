package main

import (
	"bufio"
	"fmt"
	"os"
)


type Direction string


const (
    UP          Direction = "UP"
    DOWN        Direction = "DOWN"
    LEFT        Direction = "LEFT"
    RIGHT       Direction = "RIGHT"
    UPLEFT      Direction = "UPLEFT"
    UPRIGHT     Direction = "UPRIGHT"
    DOWNLEFT    Direction = "DOWNLEFT"
    DOWNRIGHT   Direction = "DOWNRIGHT"
)


type coord2D struct {
    x   int
    y   int
}



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
    startPos := make([]coord2D, 0)
    runeGrid := make([][]rune, 0)
    for scanner.Scan() {
        for i, rune := range scanner.Text() {
            if rune == 'X' {
                startPos = append(startPos, coord2D{x: i, y: len(runeGrid)})
            }
        }

        runeGrid = append(runeGrid, []rune(scanner.Text()))
    }

    total := 0
    for _, start := range startPos {
        directions := getPossibleDirections(start, runeGrid)
        for _, direction := range directions {
            if searchDirection(start, direction, 'X', runeGrid) {
                total += 1
            }
        }
    }

    return total
}


func task02() int {
    panic("unimplemented")
}


func searchDirection(curCoord coord2D, direction Direction, nextRune rune, runeGrid [][]rune) bool {
    runeAtCoord := runeGrid[curCoord.y][curCoord.x]
    if runeAtCoord != nextRune {
        return false
    }

    if runeAtCoord == 'S' {
        return true
    }

    // Get next coordinate (current coordinate + direction)
    newCoord := curCoord
    switch direction {
        case UP:
            newCoord.y -= 1
        case DOWN:
            newCoord.y += 1
        case LEFT:
            newCoord.x -= 1
        case RIGHT:
            newCoord.x += 1
        case UPLEFT:
            newCoord.y -= 1
            newCoord.x -= 1
        case UPRIGHT:
            newCoord.y -= 1
            newCoord.x += 1
        case DOWNLEFT:
            newCoord.y += 1
            newCoord.x -= 1
        case DOWNRIGHT:
            newCoord.y += 1
            newCoord.x += 1
    }

    // Check if new coordinate is out of bounds
    if newCoord.x < 0 || newCoord.x >= len(runeGrid[0]) || newCoord.y < 0 || newCoord.y >= len(runeGrid) {
        return false
    }

    if runeAtCoord == 'X' {
        nextRune = 'M'
    } else if runeAtCoord == 'M' {
        nextRune = 'A'
    } else if runeAtCoord == 'A' {
        nextRune = 'S'
    }

    return searchDirection(newCoord, direction, nextRune, runeGrid)
}


func getPossibleDirections(start coord2D, runeGrid [][]rune) []Direction {
    canUp, canDown, canLeft, canRight := true, true, true, true
    if start.y == 0 { canUp = false }
    if start.y == len(runeGrid)-1 { canDown = false }
    if start.x == 0 { canLeft = false }
    if start.x == len(runeGrid[0])-1 { canRight = false }

    directions := make([]Direction, 0)
    if canUp &&  runeGrid[start.y - 1][start.x] == 'M' {
        directions = append(directions, UP)
    }
    if canDown && runeGrid[start.y + 1][start.x] == 'M' {
        directions = append(directions, DOWN)
    }
    if canLeft && runeGrid[start.y][start.x - 1] == 'M' {
        directions = append(directions, LEFT)
    }
    if canRight && runeGrid[start.y][start.x + 1] == 'M' {
        directions = append(directions, RIGHT)
    }
    if canLeft && canUp && runeGrid[start.y - 1][start.x - 1] == 'M' { 
        directions = append(directions, UPLEFT)
    }
    if canRight && canUp && runeGrid[start.y - 1][start.x + 1] == 'M' {
        directions = append(directions, UPRIGHT)
    }
    if canLeft && canDown && runeGrid[start.y + 1][start.x - 1] == 'M' {
        directions = append(directions, DOWNLEFT)
    }
    if canRight && canDown && runeGrid[start.y + 1][start.x + 1] == 'M' {
        directions = append(directions, DOWNRIGHT)
    }

    return directions
}
