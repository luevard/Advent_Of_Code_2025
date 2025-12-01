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
	instruction string
}
type RotateStatus int

const (
	isZero RotateStatus = iota
	haveRotated
	withoutRotation
)

func DisplayRotation(r Rotate, isRotate RotateStatus, dialValue int){
	var message string
	switch isRotate{
	case haveRotated:
		message=fmt.Sprintf(" - The dial is rotated %s%d to point at %d; during this rotation, it points at 0 once.",r.instruction,r.number,dialValue)
	default:
		message=fmt.Sprintf(" - The dial is rotated %s%d to point at %d.",r.instruction,r.number,dialValue)
	}
	println(message)
}

func RotateDial(r Rotate, dialValue int, c int) (int, int){
	rotateStatus := withoutRotation
	if(r.isRight){
		if(dialValue+r.number > 100){
			rotateStatus = haveRotated
			c++
		}
		if(dialValue==0){
			rotateStatus = isZero
			c++
		}
		dialValue=(dialValue+r.number)%100
	} else {
		if(dialValue==0){
			rotateStatus = isZero
		}
		dialValue=dialValue-r.number
		if(dialValue < 0){
			if(rotateStatus != isZero){
				rotateStatus = haveRotated
			}
			dialValue=dialValue+100
			c++
		}
	}
	DisplayRotation(r,rotateStatus,dialValue)
	return dialValue,c+r.numberOfTurn
}

func main() {
    content, err := os.Open("Puzzle_File/Day-1.txt")
    if err != nil {
        log.Fatal(err)
    }
	scanner := bufio.NewScanner(content)
	dialValue := 50
	counterZero := 0
	fmt.Printf(" - The dial starts by pointing at %d.\n",dialValue)
	for scanner.Scan(){
		line := scanner.Text()
		number, _ := strconv.Atoi(line[1:])
		rotate := Rotate{isRight: strings.Contains(string(line[0]),"R"), number: number%100, numberOfTurn: number/100,instruction: string(line[0])}
		dialValue, counterZero=RotateDial(rotate,dialValue,counterZero)
	}
	fmt.Printf("The password is: %d",counterZero)
}