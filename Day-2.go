package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func allSameStrings(a []string)(bool){
    for i := 1; i < len(a); i++ {
        if a[i] != a[0] {
            return false
        }
    }
    return true
}

func cutProductIdByN(p string, index int)([]string){
	lenProductId := len(p)
	var productIdCuted []string

	for i := 0; i < lenProductId; i += index {
		end := i + index
		if end > lenProductId {
			end = lenProductId
		}
		productIdCuted = append(productIdCuted, string(p[i:end]))
	}
	return productIdCuted
}

func checkingID(productId string)(bool){
	lenProductId := len(productId)
	for i := 1; i < lenProductId; i++{
		if(lenProductId%i==0){
			if(allSameStrings(cutProductIdByN(productId,i))){
				return true
			}
		}
	}
	return false
}

func checkingProductRange(startingRange string, endingRange string)([]int){
	var productIdCorrupt []int

	endingRangeInt, _:=strconv.Atoi(endingRange)
	startingRangeInt, _:=strconv.Atoi(startingRange)

	for i := range (endingRangeInt-startingRangeInt)+1{
		idToCheck := startingRangeInt+i
		if(checkingID(strconv.Itoa(idToCheck))){
			productIdCorrupt = append(productIdCorrupt, idToCheck)
		}
	}

	return productIdCorrupt
}

func main(){
    content, err := os.Open("Puzzle_File/Day-2.txt")

    if err != nil {
        log.Fatal(err)
    }

	scanner := bufio.NewScanner(content)
	scanner.Scan()
	line := scanner.Text()
	var productIdCorrupt = []int{}
	productionIdRange:=strings.Split(line, ",")

	for index , _ := range productionIdRange{
		rangeProduction:=strings.Split(productionIdRange[index],"-")
		for _, productCorrupt := range (checkingProductRange(rangeProduction[0],rangeProduction[1])){
			productIdCorrupt = append(productIdCorrupt, productCorrupt)
		}
	}

	total := 0 
	for index := range productIdCorrupt{
		total += productIdCorrupt[index]
	}

	println(total)
}