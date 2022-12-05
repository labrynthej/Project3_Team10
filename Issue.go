package main

var PreMemBuff = [2]int{-1, -1} // two instruction indexes
var PreALUBuff = [2]int{-1, -1} // two instruction indexes

func issue(instrArray []Instruction) {
	for i := 0; i < 4; i++ {
		if PreIssueBuff[i] != -1 { // check if each entry of PreIssueBuffer is available
			typeOfInstruction := instrArray[PreIssueBuff[i]].typeOfInstruction // get type of instruction

			switch typeOfInstruction {
			case "R":
				fallthrough
			case "I":
				fallthrough
			case "CB":
				fallthrough
			case "B":
				fallthrough
			case "IM":
				for j := 0; j < 2; j++ {
					if PreALUBuff[j] == -1 { // if empty, fill it!
						PreALUBuff[j] = PreIssueBuff[i] // move to preALU
						PreIssueBuff[i] = -1            // remove from PreALU
					}
				}
				break
			case "D":
				for j := 0; j < 2; j++ {
					if PreMemBuff[j] != -1 { // if empty, fill it!
						PreMemBuff[j] = PreIssueBuff[i] // move to PreMEM
						PreIssueBuff[i] = -1            // remove from preIssue
					}
				}
			}

			break // break out of loop after 1 entry filled so entire buffer isn't filled
		}
	}

}
