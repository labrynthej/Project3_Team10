package main

var PreMemBuff = [2]int{-1, -1} // two instruction indexes
var PreALUBuff = [2]int{-1, -1} // two instruction indexes

func issue(instrArray []Instruction) {
	counter := 0
	for i := 0; i < 4; i++ {
		if PreIssueBuff[i] != -1 {
			typeOfInstruction := instrArray[PreIssueBuff[i]].typeOfInstruction // get type of instruction

			switch typeOfInstruction {
			case "R":
			case "I":
			case "CB":
			case "B":
			case "IM":
				for j := 0; j < 2; j++ {
					if PreALUBuff[j] == -1 { // if empty, fill it!
						PreALUBuff[j] = PreIssueBuff[i] // move to preALU
						counter++
					}
				}
				break
			case "D":
				for j := 0; j < 2; j++ {
					if PreMemBuff[j] != -1 { // if empty, fill it!
						PreMemBuff[j] = PreIssueBuff[i] // move to PreMEM
						counter++
					}
				}
			}
			if counter == 2 { // move 3rd and 4th values in queue to 1st and second if 2 instructions were issued, free up 3rd and 4th
				PreIssueBuff[0] = PreIssueBuff[2]
				PreIssueBuff[2] = -1
				PreIssueBuff[1] = PreIssueBuff[3]
				PreIssueBuff[3] = -1
				break // break out of loop after 2 instructions are issued
			}
		}
	} // check next value of PreIssueBuffer
	if counter == 1 { // if only 1 instruction issued, free up first space in queue
		PreIssueBuff[0] = -1
	}
}
