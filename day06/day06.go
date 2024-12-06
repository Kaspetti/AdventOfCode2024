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

    total := 0
    dir := UP
    guardStart := guardPos

    gridWidth := len(grid[0])
    gridHeight := len(grid)

    tested := make(map[coord]map[Direction]bool)
    added := make(map[coord]bool)
    for {
        nextCoord := goForward(guardPos, dir, gridWidth, gridHeight)
        if nextCoord == guardPos {
            break
        }

        cellInFront := grid[nextCoord.y][nextCoord.x]
        if cellInFront == '#' {
            dir = (dir + 1) % 4

        }

        nextCoord = goForward(guardPos, dir, gridWidth, gridHeight)
        if nextCoord == guardPos {
            break
        }

        if _, ok := tested[guardPos]; !ok {
            tested[guardPos] = make(map[Direction]bool)
        }

        if _, ok := tested[guardPos][dir]; !ok {
            if nextCoord != guardStart {
                candidate := obstacleToTheRight(guardPos, dir, grid)
                if candidate {
                    isLoop := checkForLoop(guardPos, dir, nextCoord, grid)
                    if _, ok := added[nextCoord]; isLoop && !ok {
                        added[nextCoord] = true
                        total += 1
                    }
                }
            }
        }


        tested[guardPos][dir] = true
        guardPos = nextCoord
    }

    return total
}


func obstacleToTheRight(c coord, dir Direction, grid [][]rune) bool {
    switch dir {
        case UP:
            y := c.y
            for x := c.x + 1; x < len(grid[y]); x++ {
                if grid[y][x] == '#' {
                    return true
                }
            }
        case DOWN:
            y := c.y
            for x := c.x - 1; x >= 0; x-- {
                if grid[y][x] == '#' {
                    return true
                }
            }
        case LEFT:
            x := c.x
            for y := c.y - 1; y >= 0; y-- {
                if grid[y][x] == '#' {
                    return true
                }
            }
        case RIGHT:
            x := c.x
            for y := c.y + 1; y < len(grid); y++ {
                if grid[y][x] == '#' {
                    return true
                }
            }
    }

    return false
}


func checkForLoop(guardPos coord, dir Direction, newObstacle coord, grid [][]rune) bool {
    newGrid := make([][]rune, len(grid))
    for i := range grid {
        newGrid[i] = make([]rune, len(grid[i]))
        copy(newGrid[i], grid[i])
    }
    newGrid[newObstacle.y][newObstacle.x] = '#'

    visited := make(map[coord]map[Direction]bool)
    gridWidth := len(grid[0])
    gridHeight := len(grid)

    for {
        if _, ok := visited[guardPos][dir]; ok {
            return true
        }

        if _, ok := visited[guardPos]; !ok {
            visited[guardPos] = make(map[Direction]bool)
        }

        visited[guardPos][dir] = true

        nextCoord := goForward(guardPos, dir, gridWidth, gridHeight)
        if nextCoord == guardPos {
            break
        }
        
        cellInFront := newGrid[nextCoord.y][nextCoord.x]
        if cellInFront == '#' {
            dir = (dir + 1) % 4
        }

        nextCoord = goForward(guardPos, dir, gridWidth, gridHeight)
        if nextCoord == guardPos {
            break
        }

        guardPos = nextCoord
    }

    return false
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
