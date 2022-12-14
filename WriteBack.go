package main

func writeBack() {
	// postALUBuff = [2]int{-1, -1} first number is value, second is instr index

	aluVal := postALUBuff[0]
	aluIndex := postALUBuff[1]

	registerMap[uint8(aluIndex)] = aluVal

	postALUBuff[0] = -1
	postALUBuff[1] = -1

	// postMemBuff = [2]int{-1, -1} first number is value, second is instr index
	memVal := postMemBuff[0]
	memIndex := postMemBuff[1]

	registerMap[uint8(memIndex)] = memVal
}
