package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"sort"
	"strings"
    "fmt"
)

type Range struct {
    startingRange *big.Int
    endingRange   *big.Int
}

func countUniqueIDs(ranges []Range) *big.Int {
    if len(ranges) == 0 {
        return big.NewInt(0)
    }

    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i].startingRange.Cmp(ranges[j].startingRange) < 0
    })

    merged := []Range{}
    current := ranges[0]

    for _, r := range ranges[1:] {
        if r.startingRange.Cmp(new(big.Int).Add(current.endingRange, big.NewInt(1))) <= 0 {
            if r.endingRange.Cmp(current.endingRange) > 0 {
                current.endingRange = r.endingRange
            }
        } else {
            merged = append(merged, current)
            current = r
        }
    }
    merged = append(merged, current)

    total := big.NewInt(0)
    for _, r := range merged {
        size := new(big.Int).Sub(r.endingRange, r.startingRange)
        size.Add(size, big.NewInt(1))
        total.Add(total, size)
    }

    return total
}


func main() {
    content, err := os.Open("Puzzle_File/Day-5.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(content)
    ranges := []Range{}

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "-")
        if len(parts) == 2 {
            startingRange := new(big.Int)
            endingRange := new(big.Int)
            startingRange.SetString(parts[0], 10)
            endingRange.SetString(parts[1], 10)

            ranges = append(ranges, Range{startingRange: startingRange, endingRange: endingRange})
        }
    }
    total := countUniqueIDs(ranges)
    fmt.Println(total)
}
