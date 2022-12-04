package main

func controlUnit(instrArray []Instruction) {
	// count := 96
	cycle := 1
	for i, _ := range instrArray {
		cycle = i
		runALU(instrArray)
		runMem(instrArray)
		writeBack() // last instruction

	}
	cycle++
}
