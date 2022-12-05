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
		writeBack() // last instruction
		runALU(instrArray)
		runMem(instrArray)
		issue(instrArray)
		fetchInstr(instrArray, count)

		printPipeline(instrArray, f, cycle)

		count = count + 4
	}
}
