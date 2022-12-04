package main

var PreIssueBuff = [4]int{-1, -1, -1, -1} // four indexes

func fetchInstr() {

}

func fetchStr(idx int) string {

	return memoryMap[idx]
}
