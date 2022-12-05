package main

import (
	"fmt"
	"os"
)

var count = 96

func controlUnit(instrArray []Instruction) {
	f, fileErr := os.Create("test.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	cycle := 1
	for i := range instrArray {
		cycle = i
		issue(instrArray)
		fetchInstr(instrArray, count)
		runALU(instrArray)
		runMem(instrArray)
		writeBack() // last instruction
		printPipeline(instrArray, f)

		count = count + (i * 8)
	}
	cycle++
}
