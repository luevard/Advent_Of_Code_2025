package main

import (
    "log"
    "os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

type Rotate struct{
	isRight bool
	number int
	numberOfTurn int
}

func RotateDial(r Rotate, dialValue int, c int) (int, int){
	if(r.isRight){
		if(dialValue+r.number > 100 && dialValue!=0){
			fmt.Println("Rotation R")
			c++
		}
		dialValue=(dialValue+r.number)%100
		fmt.Println("The dial is rotated R",r.number," to point at",dialValue)
	} else {
		if(dialValue == 0){
			dialValue=dialValue+100
		}
		dialValue=dialValue-r.number
		if(dialValue < 0){
			fmt.Println("Rotation L")
			dialValue=dialValue+100
			c++
		}
		fmt.Println("The dial is rotated L",r.number," to point at",dialValue)
	}
	if(dialValue == 0){
		c++
	}
	return dialValue,c+r.numberOfTurn
}

func main() {
    content, err := os.Open("Puzzle_File/Day-1.txt")
    if err != nil {
        log.Fatal(err)
    }
	scanner := bufio.NewScanner(content)
	rotateArray := make([]Rotate, 1)

	for scanner.Scan(){
		line := scanner.Text()
		number, _ := strconv.Atoi(line[1:])
		rotate := Rotate{isRight: strings.Contains(string(line[0]),"R"), number: number%100, numberOfTurn: number/100}
		rotateArray = append(rotateArray, rotate)
	}

	dialValue := 50
	counterZero := 0

	for _,rotate := range rotateArray{
		dialValue, counterZero=RotateDial(rotate,dialValue,counterZero)
	}
	println(counterZero)
}