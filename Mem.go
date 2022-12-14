package main

var postMemBuff = [2]int{-1, -1} //first number is value, second is instr index

func runMem(instrArray []Instruction) {
	for i := 0; i < 2; i++ {
		if PreMemBuff[i] != -1 {
			currentInstr := instrArray[PreMemBuff[i]]
			switch currentInstr.op {
			// D format instructions
			case "LDUR":
				postMemBuff[0] = dataSlice[registerMap[currentInstr.rn]+int(currentInstr.address)*4]
				break
			case "STUR":
				dataSlice[registerMap[currentInstr.rn]+int(currentInstr.address)*4] = registerMap[currentInstr.rt]
				break
				// IM format instructions might go here too
			}
			postMemBuff[1] = PreMemBuff[0]
			PreMemBuff[0] = PreMemBuff[1]
			PreMemBuff[1] = -1
		}
	}

}
