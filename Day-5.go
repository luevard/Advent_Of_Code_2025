package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
	"strings"
)

type Range struct{
	startingRange []int
	endingRange []int
}

func splitInt(valueToSplit *big.Int)([]int){
	str := valueToSplit.String()
	arrayTab := make([]int, len(str))
	for i, int32Unicode := range str {
		arrayTab[i] = int(int32Unicode - '0')
	}
	return arrayTab
}

func isBiggerOrEqual(a1 []int, a2 []int) bool {
    if len(a1) > len(a2) {
        return true
    } else if len(a1) < len(a2) {
        return false
    }

	for i := 0; i < len(a1); i++ {
        if a1[i] > a2[i] {
            return true
        }
        if a1[i] < a2[i] {
            return false
        }
    }

    return true
}


func checkId(ranges []Range, freshId []int)(bool){
	for rang := range len(ranges){
		if(isBiggerOrEqual(freshId,ranges[rang].startingRange) && isBiggerOrEqual(ranges[rang].endingRange,freshId)){
			return true
		}
	}
	return false
}


func checkIds(ranges []Range, freshIds [][]int)(int){
	cpt:=0
	for freshId := range len(freshIds){
		if(checkId(ranges,freshIds[freshId])){
			cpt++
		}
	}
	return cpt
}

func main(){
    content, err := os.Open("Puzzle_File/Day-5.txt")
    if err != nil {
        log.Fatal(err)
    }
	scanner := bufio.NewScanner(content)
	ranges:=[]Range{}
	freshIds:=[][]int{}
	for scanner.Scan(){
		text:=scanner.Text()
		textArray:=strings.Split(text,"-")
		if(len(textArray) == 2){
			startingRange := new(big.Int)
			endingRange := new(big.Int)
			startingRange.SetString(textArray[0], 10)
			endingRange.SetString(textArray[1], 10)
			ranges = append(ranges, Range{startingRange:splitInt(startingRange),endingRange:splitInt(endingRange)})
		}else{
			freshId := new(big.Int)
			freshId.SetString(text, 10)
			freshIds = append(freshIds, splitInt(freshId))
		}
	}

	total:=checkIds(ranges,freshIds)
	println(total)
}
