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
	cycle := 0
	for i := range instrArray {
		cycle = i + 1
		issue(instrArray)
		fetchInstr(instrArray, count)
		runALU(instrArray)
		runMem(instrArray)
		writeBack() // last instruction
		printPipeline(instrArray, f)

		count = count + 8
	}
	cycle++
}
