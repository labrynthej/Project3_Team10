package main

var postMemBuff = [2]int{-1, -1} //first number is value, second is instr index

func memToWriteBack() {
	postALUBuff = [2]int{1100, 9}
	writeBack()
}
