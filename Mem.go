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
		// IM format instructions
		case "MOVZ":
			registerMap[currentInstr.rd] = 0
			registerMap[currentInstr.rd] = int(currentInstr.field<<(currentInstr.shamt*16)) &
				(0xFFFFFFFF << (currentInstr.shamt * 16))
			break
		case "MOVK":
			registerMap[currentInstr.rd] = registerMap[currentInstr.rd] +
				int(currentInstr.field<<(currentInstr.shamt*16))
			break
		}
		PreMemBuff[0] = PreMemBuff[1]
		PreMemBuff[1] = -1
	}
}
