package main

import (
	"fmt"
	"strconv"
)

func intInArray(search int, array []int) bool {
	for _, v := range array {
		if search == v {
			return true
		}
	}
	return false
}

func main() {
	var input string
	for {
		fmt.Print(">> ")
		fmt.Scanln(&input)
		sumaSimbolo := "+"
		restaSimbolo := "-"
		leninput := len(input)
		simbolosPos := make([]int, 0)
		sumasPos := make([]int, 0)
		restasPos := make([]int, 0)
		for pos, ch := range input {
			char := string(ch)
			if char == sumaSimbolo {
				sumasPos = append(sumasPos, pos)
				simbolosPos = append(simbolosPos, pos)
			} else if char == restaSimbolo {
				restasPos = append(restasPos, pos)
				simbolosPos = append(simbolosPos, pos)
			}
		}
		maxRange := len(simbolosPos) - 1
		resultado := 0
		for indx, _ := range simbolosPos { // 1,5
			if indx == 0 {
				num, _ := strconv.Atoi(input[0:simbolosPos[indx]])
				// fmt.Printf("Parsing0: %d\n", num)
				resultado += num
			}
			if indx < maxRange {
				num, _ := strconv.Atoi(input[simbolosPos[indx]+1 : simbolosPos[indx+1]])
				// fmt.Printf("Parsing1: %d\n", num)
				if intInArray(simbolosPos[indx], sumasPos) {
					resultado += num
				} else if intInArray(simbolosPos[indx], restasPos) {
					resultado -= num
				}
			} else {
				num, _ := strconv.Atoi(input[simbolosPos[indx]+1 : leninput])
				// fmt.Printf("Parsingf: %d\n", num)
				if intInArray(simbolosPos[indx], sumasPos) {
					resultado += num
				} else if intInArray(simbolosPos[indx], restasPos) {
					resultado -= num
				}
			}
		}
		if len(simbolosPos) < 1 {
			fmt.Println(input)
		} else {
			fmt.Println(resultado)
		}
		if input == "0" {
			break
		}
	}
}
