package main

var postALUBuff = [2]int{-1, -1} //first number is value, second is instr index

func runALU(instrArray []Instruction) {
	if PreALUBuff[0] != -1 {
		currentInstr := instrArray[PreALUBuff[0]]
		switch currentInstr.op {
		// R format instructions
		case "SUB": // 	rd = rn - rm
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] - registerMap[currentInstr.rm]
			break
		case "AND": // rd = rm & rn
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] & registerMap[currentInstr.rm]
			break
		case "ADD": // rd = rm + rn
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] + registerMap[currentInstr.rm]
			break
		case "ORR": // rd = rm | rn
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] | registerMap[currentInstr.rm]
			break
		case "EOR": // rd = rm ^ rn
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] ^ registerMap[currentInstr.rm]
			break
		case "LSR": // rn shifted shamt
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] >> registerMap[currentInstr.shamt]
			break
		case "LSL": // rd = rn << shamt
			registerMap[currentInstr.rd] = (registerMap[currentInstr.rn]) << registerMap[currentInstr.shamt]
			break
		case "ASR": // rd = rn >> shamt pad with sign bit
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] >> registerMap[currentInstr.shamt]
			break

		// I format instructions
		case "ADDI": // rd = rn + im
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] + int(currentInstr.im)
			break
		case "SUBI": // rd = rn - im
			registerMap[currentInstr.rd] = registerMap[currentInstr.rn] - int(currentInstr.im)
			break

		// B and CB format instructions
		case "B": // PC = PC +- (4 * offset)
			count = int(currentInstr.offset)
			break
		case "CBZ": // if (conditional == 0) {PC = 4 * offset}
			if currentInstr.conditional == 0 {
				count = int(currentInstr.offset)
			}
			break
		case "CBNZ": // if (conditional == 1) {PC = 4 * offset}
			if currentInstr.conditional != 0 {
				count = int(currentInstr.offset)
			}
			break
		}
	}
	postALUBuff[0] = postALUBuff[1]
	postALUBuff[1] = -1
}
