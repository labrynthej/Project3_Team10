package main

var postALUBuff = [2]int{-1, -1} //first number is value, second is instr index

func runALU(instrArray []Instruction) {
	for i := 0; i < 2; i++ {
		if PreALUBuff[i] != -1 {
			currentInstr := instrArray[PreALUBuff[i]]
			switch currentInstr.op {
			// R format instructions
			case "SUB": // 	rd = rn - rm
				postALUBuff[0] = registerMap[currentInstr.rn] - registerMap[currentInstr.rm]
				break
			case "AND": // rd = rm & rn
				postALUBuff[0] = registerMap[currentInstr.rn] & registerMap[currentInstr.rm]
				break
			case "ADD": // rd = rm + rn
				postALUBuff[0] = registerMap[currentInstr.rn] + registerMap[currentInstr.rm]
				break
			case "ORR": // rd = rm | rn
				postALUBuff[0] = registerMap[currentInstr.rn] | registerMap[currentInstr.rm]
				break
			case "EOR": // rd = rm ^ rn
				postALUBuff[0] = registerMap[currentInstr.rn] ^ registerMap[currentInstr.rm]
				break
			case "LSR": // rn shifted shamt
				postALUBuff[0] = registerMap[currentInstr.rn] >> registerMap[currentInstr.shamt]
				break
			case "LSL": // rd = rn << shamt
				postALUBuff[0] = (registerMap[currentInstr.rn]) << registerMap[currentInstr.shamt]
				break
			case "ASR": // rd = rn >> shamt pad with sign bit
				postALUBuff[0] = registerMap[currentInstr.rn] >> registerMap[currentInstr.shamt]
				break

			// I format instructions
			case "ADDI": // rd = rn + im
				postALUBuff[0] = registerMap[currentInstr.rn] + int(currentInstr.im)
				break
			case "SUBI": // rd = rn - im
				postALUBuff[0] = registerMap[currentInstr.rn] - int(currentInstr.im)
				break
			}
			postALUBuff[1] = int(currentInstr.rd)
			PreALUBuff[i] = -1
			break
		}
	}

}
