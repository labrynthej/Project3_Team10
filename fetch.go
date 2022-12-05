package main

import (
	"fmt"
	"strconv"
)

var PreIssueBuff = [4]int{-1, -1, -1, -1} // four indexes

func fetchInstr(instrArray []Instruction, count int) {
	j := 0
	for i := 0; i < 3; i++ {
		if PreIssueBuff[i] == -1 {
			j++
			if j < 3 {
				PreIssueBuff[i] = readFromCache(count)
			}
			count = count + 4
		}
	}
	fmt.Println(PreIssueBuff)

}

func setValue(instr []Instruction, i int) {
	lineValue, _ := strconv.ParseUint(instr[i].rawInstruction, 2, 32)

	if lineValue > 335544320 || lineValue == 0 {
		// assign lineValue and 11 bit opcode for setting the instruction
		instr[i].lineValue = lineValue
		instr[i].opcode = lineValue >> 21

		// set values for instruction type "R" | opcode | Rm | Shamt | Rn | Rd |
		if instr[i].typeOfInstruction == "R" {
			instr[i].src1Reg = int((lineValue & 0x3E0) >> 5)
			instr[i].src2Reg = int((lineValue & 0x1F0000) >> 16)
			instr[i].destReg = int(lineValue & 0x1F)
			instr[i].shamt = uint8((lineValue & 0xFC00) >> 11)
		}

		// set values for instruction type "D" | opcode | address | op2 | Rn | Rt |
		if instr[i].typeOfInstruction == "D" {
			instr[i].src1Reg = int((lineValue & 0x3E0) >> 5)
			instr[i].address = uint8((lineValue & 0x1FF000) >> 12)
			instr[i].op2 = uint8((lineValue & 0xC00) >> 10)
			instr[i].destReg = int(lineValue & 0x1F)
		}

		// set values for instruction type "I" | opcode | immediate | Rn | Rd |
		if instr[i].typeOfInstruction == "I" {
			instr[i].opcode = lineValue >> 22
			instr[i].src1Reg = int((lineValue & 0x3E0) >> 5)
			instr[i].im = uint8(signedVariable(lineValue&0x3FFC00>>10, 12))
			instr[i].destReg = int(lineValue & 0x1F)
		}

		// set values for instruction type "B" | opcode | offset |
		if instr[i].typeOfInstruction == "B" {
			instr[i].opcode = lineValue >> 26
			instr[i].offset = signedVariable(lineValue&0x3FFFFFF, 26)
		}

		// set values for instruction type "CB" (conditional B) | opcode | offset |
		if instr[i].typeOfInstruction == "CB" {
			instr[i].opcode = lineValue >> 24
			instr[i].offset = signedVariable(lineValue&0xFFFFE0>>5, 19)
			instr[i].conditional = uint8(lineValue & 0x1F)
		}

		// set values for instruction type "IM" | opcode | shift | field | Rd |
		if instr[i].typeOfInstruction == "IM" {
			instr[i].opcode = lineValue >> 23
			instr[i].shamt = uint8(lineValue & 0x600000 >> 21)
			instr[i].field = lineValue & 0x1FFFE0 >> 5
			instr[i].destReg = int(lineValue & 0x1F)
		}

		if instr[i].op == "BREAK" {
		}
	}
}

func fetchStr(instr []Instruction, idx int) string {
	if idx != -1 {
		return fmt.Sprintf("%s", instr[idx].op)
	}

	return ""
}
