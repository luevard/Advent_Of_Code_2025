package main

import (
	"bufio"
	"log"
	"os"
)

func convertGridOfCommercial(s string)([]int){
	// 64 = @
	// 46 = .
	line := []int{} 	
	for b := range len(s){
		if(s[b]==64){
			line = append(line, 1)
		}else{
			line = append(line, 0)
		}
	}
	return line
}

func RemovePapers(grid [][]int)([][]int){
	maxLine := len(grid)
	if(maxLine==0){
		return nil
	}
	itemMax := len(grid[maxLine-1])
	for line := range maxLine{
		for item := range itemMax{
			if(grid[line][item]==2){
				grid[line][item]=0
			}
		}
	}
	return grid
}

func checkPapers(grid [][]int)([][]int,int){
	maxLine:=len(grid)
	if(maxLine==0){
		return nil,0
	}
	itemMax := len(grid[maxLine-1])
	total:=0

	for line := range maxLine{
		for item := range itemMax{
			if(grid[line][item]==1){
				cpt:=0
				if(line>0){
					if(grid[line-1][item]>=1){
						cpt++
					}
					if(item+1 < itemMax){
						if(grid[line-1][item+1]>=1){
							cpt++
						}
					}
					if(item > 0){
						if(grid[line-1][item-1]>=1){
							cpt++
						}
					}
				}
				if(item>0){
					if(grid[line][item-1]>=1){
						cpt++
					}
					if(line+1 < maxLine){
						if(grid[line+1][item-1]>=1){
							cpt++
						}
					}
				}
				if(item+1 < itemMax){
					if(grid[line][item+1]>=1){
						cpt++
					}
				}
				if(line+1 < maxLine){
					if(grid[line+1][item]>=1){
						cpt++
					}
				}
				if(item+1 < itemMax && line+1 < maxLine){
					if(grid[line+1][item+1]>=1){
						cpt++
					}
				}
				if(cpt<4){
					total++
					grid[line][item]= 2
				}
			}
		}
	}
	return grid,total
}

func main(){
    content, err := os.Open("Puzzle_File/Day-4.txt")
    if err != nil {
        log.Fatal(err)
    }
	scanner := bufio.NewScanner(content)
	grid := [][]int{}
	result := 0
	total := -1
	line:=[]int{}
	for scanner.Scan(){
		text:=scanner.Text()
		line=convertGridOfCommercial(text)
		grid = append(grid, line)
	}

	for total != 0 {
		grid, total = checkPapers(RemovePapers(grid))
		result=result+total
	}

	println(result)
}