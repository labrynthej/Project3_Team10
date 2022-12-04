package main

import (
	"fmt"
	"os"
)

func controlUnit(instrArray []Instruction) {
	f, fileErr := os.Create("test.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	count := 96
	cycle := 1
	for i, _ := range instrArray {
		cycle = i
		fetchInstr(instrArray, count)
		issue(instrArray)
		runALU(instrArray)
		runMem(instrArray)
		writeBack() // last instruction
		printPipeline(instrArray, f)

		count = count + (i * 8)
	}
	cycle++
}
