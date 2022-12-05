package main

var postMemBuff = [2]int{-1, -1} //first number is value, second is instr index

func runMem(instrArray []Instruction) {
	if PreMemBuff[0] != -1 {
		currentInstr := instrArray[PreMemBuff[0]]
		switch currentInstr.op {
		// D format instructions
		case "LDUR":
			registerMap[currentInstr.rt] = dataSlice[registerMap[currentInstr.rn]+int(currentInstr.address)*4]
			break
		case "STUR":
			dataSlice[registerMap[currentInstr.rn]+int(currentInstr.address)*4] = registerMap[currentInstr.rt]
			break
			// IM format instructions might go here too
		}
		PreMemBuff[0] = PreMemBuff[1]
		PreMemBuff[1] = -1
	}
}
