package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/snapshot-chromedp/render"
)


func main() {
    task01()
}


func task01() {
    f, err := os.Open("input")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(f)

    col0, col1 := make([]int, 0), make([]int, 0)
    for scanner.Scan() {
        line := scanner.Text()
        cols := strings.Split(line, "   ")
        
        num0, err := strconv.Atoi(cols[0])
        if err != nil {
            panic(err)
        }
        num1, err := strconv.Atoi(cols[1])
        if err != nil {
            panic(err)
        }

        col0 = append(col0, num0)
        col1 = append(col1, num1)
    }

    sort.Slice(col0, func(i, j int) bool {
        return col0[i] < col0[j]
    })
    sort.Slice(col1, func(i, j int) bool {
        return col1[i] < col1[j]
    })

    totalDistance := 0
    distances := make([]int, 0)
    minDistance := 100000
    maxDistance := -1
    for i := range col0 {
        distance := int(math.Abs(float64(col0[i] - col1[i])))
        totalDistance += distance
        distances = append(distances, distance)

        minDistance = min(minDistance, distance)
        maxDistance = max(maxDistance, distance)
    }

    // Visualization
    distanceRange := maxDistance - minDistance
    binAmount := 10
    binSize := distanceRange / (binAmount - 1)

    bins := make([]int, binAmount)
    for _, distance := range distances {
        bins[distance / binSize] += 1
    }

    barBins := make([]string, binAmount)
    for i := range barBins {
        barBins[i] = fmt.Sprintf("%d-%d", i * binSize, (i + 1)*binSize - 1)
    }
    barValues := make([]opts.BarData, 10)
    for i, bin := range bins {
        barValues[i] = opts.BarData{Value: bin}
    }

    bar := charts.NewBar()
    bar.SetGlobalOptions(
        charts.WithTitleOpts(opts.Title{
            Title: "Distances in range",
        }),
        charts.WithXAxisOpts(opts.XAxis{
            AxisLabel: &opts.AxisLabel{
                Interval: "",
                Rotate: 45,
            },
        }),
        charts.WithAnimation(false),
    )

    bar.SetXAxis(barBins).
        AddSeries("Distances", barValues)

    render.MakeChartSnapshot(bar.RenderContent(), "task01.png")


    fmt.Println(totalDistance)
}
