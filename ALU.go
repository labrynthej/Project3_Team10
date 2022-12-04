package main

var postALUBuff = [2]int{-1, -1} //first number is value, second is instr index

func toWriteBack() {
	postALUBuff = [2]int{12, 3}
	writeBack()
}

func aluValue() int {

	return 0
}
