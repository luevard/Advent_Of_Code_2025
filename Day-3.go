package main

import (
	"bufio"
	"log"
	"os"
	"math/big"
	"math"
)

func splitInt(valueToSplit *big.Int)([]int){
	str := valueToSplit.String()
	arrayTab := make([]int, len(str))
	for i, int32Unicode := range str {
		arrayTab[i] = int(int32Unicode - '0')
	}
	return arrayTab
}

func concatIntArray(valueToConcat []int)(int){
	result := 0
	for i := 0; i < len(valueToConcat); i++ {
		power := len(valueToConcat) - 1 - i
		result += valueToConcat[i] * int(math.Pow(10, float64(power)))
	}
	return result
}

func check(n []int)([]int){
	tab := make([]int, 0, 12)
	start := 0
	remaining := 12

	for remaining > 0 {
		end := len(n) - remaining
		maxDigit := -1
		maxIndex := start

		for i := start; i <= end; i++ {
			if n[i] > maxDigit {
				maxDigit = n[i]
				maxIndex = i
			}
		}

		tab = append(tab, maxDigit)
		start = maxIndex + 1
		remaining--
	}
	return tab
}


func main() {
    content, err := os.Open("Puzzle_File/Day-3.txt")
    if err != nil {
        log.Fatal(err)
    }
	scanner := bufio.NewScanner(content)
	total := 0
	for scanner.Scan(){
		value := new(big.Int)
		value.SetString(scanner.Text(), 10)
		total = total + concatIntArray(check(splitInt(value)))
	}
	println(total)
}