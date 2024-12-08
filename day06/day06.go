package main

import (
	"bufio"
	"fmt"
	"os"
)


type Direction int
const (
    UP      Direction = iota
    RIGHT
    DOWN
    LEFT
)


type coord struct {
    x   int
    y   int
}


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

    grid := make([][]rune, 0)
    guardPos := coord{-1, -1}

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := []rune(scanner.Text())
        grid = append(grid, line)

        if (guardPos != coord{-1, -1}) {
            continue
        }

        for x, r := range line {
            if r == '^' {
                guardPos = coord{x, len(grid)-1}
            }
        }
    }

    visited := make(map[coord]bool)
    direction := UP

    gridWidth := len(grid[0])
    gridHeight := len(grid)
    // While the guard position is within the grid
    for {
        visited[guardPos] = true

        nextCoord := goForward(guardPos, direction, gridWidth, gridHeight)
        if nextCoord == guardPos {
            break
        }
        
        cellInFront := grid[nextCoord.y][nextCoord.x]
        if cellInFront == '#' {
            direction = (direction + 1) % 4
        }

        nextCoord = goForward(guardPos, direction, gridWidth, gridHeight)
        if nextCoord == guardPos {
            break
        }

        guardPos = nextCoord
    }

    return len(visited)
}


func task02() int {
    f, err := os.Open("sampleinput")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    grid := make([][]rune, 0)
    guardPos := coord{-1, -1}

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := []rune(scanner.Text())
        grid = append(grid, line)

        if (guardPos != coord{-1, -1}) {
            continue
        }

        for x, r := range line {
            if r == '^' {
                guardPos = coord{x, len(grid)-1}
            }
        }
    }

    gridWidth := len(grid[0])
    gridHeight := len(grid)
    total := 0

    // Bruteforce check
    for y := 0; y < gridHeight; y++ {
        for x := 0; x < gridWidth; x++ {
            if grid[y][x] != '.' {
                continue
            }

            grid[y][x] = '#'
            if isLoop(guardPos, UP, grid) {
                total += 1
            }

            grid[y][x] = '.'
        }
    }

    return 0
}


func isLoop(guardPos coord, direction Direction, grid [][]rune) bool {
    visited := make(map[coord]bool)
    
    gridWidth := len(grid[0])
    gridHeight := len(grid)

    for {
        visited[guardPos] = true

        newCoord := goForward(guardPos, direction, gridWidth, gridHeight)        
        // We left the grid == not in loop
        if newCoord == guardPos {
            return false
        }

        cellInFront := grid[newCoord.y][newCoord.x]
        if cellInFront == '#' {
            if newCoord.x == 3 && newCoord.y == 6 {
                fmt.Println("hai")
            }
            direction = (direction + 1) % 4
        }

        newCoord = goForward(guardPos, direction, gridWidth, gridHeight)
        if newCoord == guardPos {
            return false
        }

        if visited[guardPos] && visited[newCoord] {
            return true
        }

        guardPos = newCoord
    }
}


func goForward(c coord, dir Direction, gridWidth, gridHeight int) coord {
    var nextCoord coord
    switch dir {
        case UP:
            nextCoord = coord{c.x, c.y - 1}
            if nextCoord.y < 0 { return c }
        case DOWN:
            nextCoord = coord{c.x, c.y + 1}
            if nextCoord.y >= gridHeight  { return c }
        case LEFT:
            nextCoord = coord{c.x - 1, c.y}
            if nextCoord.x < 0 { return c }
        case RIGHT:
            nextCoord = coord{c.x + 1, c.y}
            if nextCoord.x >= gridWidth { return c }
    }

    return nextCoord
}
