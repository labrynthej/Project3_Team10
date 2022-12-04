package main

import "fmt"

var PreIssueBuff = [4]int{-1, -1, -1, -1} // four indexes

func fetchInstr(instrArray []Instruction, count int) {

	for i := 0; i < 3; i++ {
		if PreIssueBuff[i] == -1 {
			PreIssueBuff[i], PreIssueBuff[i+1] = readFromCache(count)
			break
		}
	}

}

func fetchStr(instr []Instruction, idx int) string {
	if idx != -1 {
		return fmt.Sprintf("%s", instr[idx].op)
	}

	return ""
}
