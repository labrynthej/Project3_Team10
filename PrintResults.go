package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func printResults(instrArray []Instruction, fileName string) {

	file, fileErr := os.Create(fileName)
	if fileErr != nil {
		fmt.Println(fileErr)
	}
	i := 0
	count := 0
	for instrArray[i].typeOfInstruction != "BREAK" { // loop through each array of structs
		switch instrArray[i].typeOfInstruction {
		// print results for R Type instructions == opcode (11 bits), Rm (5 bits), Shamt (6 bits), Rn (5 bits), Rd (5 bits)
		case "R":
			// print separated binary opcode
			for j := 0; j < 32; j++ {
				if j == 11 || j == 16 || j == 22 || j == 27 { // print spaces to separate
					_, _ = file.WriteString(" ")
				}
				if j <= 10 { // print binary for opcode
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 11 && j <= 15 { // print binary for Rm/R2
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 16 && j <= 21 { // print binary for Shamt
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 22 && j <= 26 { // print binary for Rn/R1
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 27 && j <= 31 { // print binary for Rd/R3
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				}
			}
			_, _ = file.WriteString(" " + strconv.Itoa(instrArray[i].programCnt) + " " + instrArray[i].op + " R" +
				strconv.Itoa(int(instrArray[i].rd)) + ", R" + strconv.Itoa(int(instrArray[i].rn)) +
				", R" + strconv.Itoa(int(instrArray[i].rm))) // print pc, type, Rm, Shamt, Rn, Rd
			break
		// print results for D type instruction == opcode (11 bits), address (9 bits), op2 (2 bits), Rn (5 bits), Rt (5 bits)
		case "D":
			// print seperated binary opcode
			for j := 0; j < 32; j++ {
				if j == 11 || j == 20 || j == 22 || j == 27 {
					_, _ = file.WriteString(" ")
				}
				if j <= 10 { // print binary for opcode
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 11 && j <= 20 { // print binary for address
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 21 && j <= 22 { // print binary for op2
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 23 && j <= 27 { // print binary for Rn
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 28 && j <= 31 { // print binary for Rt
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				}
			}
			_, _ = file.WriteString(" " + strconv.Itoa(instrArray[i].programCnt) + " " + instrArray[i].op + " R" +
				strconv.Itoa(int(instrArray[i].rt)) + ", [R" + strconv.Itoa(int(instrArray[i].rn)) + ", #" +
				strconv.Itoa(int(instrArray[i].address)) + "]") // print pc, type, Rt, Rn, address
			break
		// print results for I type instruction == opcode (10 bits), immediate (12 bits), Rn (5 bits), Rd (5 bits)
		case "I":
			// print separated binary opcode
			for j := 0; j < 32; j++ {
				if j == 10 || j == 22 || j == 27 {
					_, _ = file.WriteString(" ")
				}
				if j <= 9 { // print binary for opcode
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 10 && j <= 21 { // print binary for immediate
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 22 && j <= 26 { // print binary for Rn
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 27 && j <= 31 { // print binary for Rd
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				}
			}
			_, _ = file.WriteString(" " + strconv.Itoa(instrArray[i].programCnt) + " " + instrArray[i].op + " R" +
				strconv.Itoa(int(instrArray[i].rd)) + ", R" + strconv.Itoa(int(instrArray[i].rn)) +
				", #" + strconv.Itoa(int(instrArray[i].im))) // print pc, type, Rd, rn, im
			break
		// print results for B type instruction == opcode (6 bits), offset (26 bits)
		case "B":
			// print separated binary code
			for j := 0; j < 32; j++ {
				if j == 6 {
					_, _ = file.WriteString(" ")
				}
				if j <= 6 { // print binary for opcode
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 7 && j <= 31 { // print binary for offset
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				}
			}
			_, _ = file.WriteString(" " + strconv.Itoa(instrArray[i].programCnt) + " " + instrArray[i].op + " #" +
				strconv.Itoa(int(instrArray[i].offset))) // print pc, type, offset
			break
		// print results for CB type instructions == opcode (8 bits), offset (19 bits), conditional (5 bits)
		case "CB":
			// print separated binary code
			for j := 0; j < 32; j++ {
				if j == 8 || j == 27 {
					_, _ = file.WriteString(" ")
				}
				if j <= 8 { // print binary for opcode
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 8 && j <= 26 { // print binary for offset
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 27 && j <= 31 { // print binary for conditional
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				}
			}
			_, _ = file.WriteString(" " + strconv.Itoa(instrArray[i].programCnt) + " " + instrArray[i].op + " R" +
				strconv.Itoa(int(instrArray[i].conditional)) + ", " + strconv.Itoa(int(instrArray[i].offset)))
			break
		// print results for IM type instructions == opcode (9 bits), shift code (2 bits), field (16 bits), Rd (5 bits)
		case "IM":
			// print seperated binary code
			for j := 0; j < 32; j++ {
				if j == 9 || j == 11 || j == 27 {
					_, _ = file.WriteString(" ")
				}
				if j <= 9 { // print binary for opcode
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 10 && j <= 11 { // print binary for shift code
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 12 && j <= 27 { // print binary for field
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				} else if j >= 28 && j <= 31 {
					_, _ = file.WriteString(string(instrArray[i].rawInstruction[j]))
				}
			}
			_, _ = file.WriteString(" " + strconv.Itoa(instrArray[i].programCnt) + " " + instrArray[i].op + " R" +
				strconv.Itoa(int(instrArray[i].rd)) + ", " + strconv.Itoa(int(instrArray[i].field)) + ", LSL " +
				strconv.Itoa(int(instrArray[i].shamt)))
			break
		case "N/A":
			_, _ = file.WriteString(instrArray[i].rawInstruction + " " +
				strconv.Itoa(instrArray[i].programCnt) + " " + "NOP")
			break
		default:
			println("Invalid value on line " + strconv.Itoa(i))
		}

		_, _ = file.WriteString("\n")
		i++
	}
	_, _ = file.WriteString(instrArray[i].rawInstruction + " " + strconv.Itoa(instrArray[i].programCnt) + " BREAK\n")
	for i = i + 1; i < len(instrArray); i++ {

		lineValue, _ := strconv.ParseUint(instrArray[i].rawInstruction, 2, 32)
		count--
		_, _ = file.WriteString(instrArray[i].rawInstruction + " " + strconv.Itoa(instrArray[i].programCnt) +
			" " + strconv.Itoa(int(signedVariable(lineValue, 32))) + "\n")
		dataSlice[instrArray[i].programCnt] = int(signedVariable(lineValue, 32))
	}
}

func printSimulation(sim Instruction, f *os.File) {

	fmt.Fprintln(f, "====================")
	fmt.Fprintf(f, "Cycle:%d\t%d\t%s\n", sim.cycle, sim.programCnt, instructionString(sim))

	// print current register
	fmt.Fprint(f, "\nRegisters:\n")
	fmt.Fprintf(f, "r00:\t%s", mapToString(registerMap, 8))
	fmt.Fprintf(f, "\nr08:\t%s", mapToString(registerMap, 16))
	fmt.Fprintf(f, "\nr16:\t%s", mapToString(registerMap, 24))
	fmt.Fprintf(f, "\nr24:\t%s\n", mapToString(registerMap, 32))

	// print data
	fmt.Fprintf(f, "\nData:")
	var keys []int
	max := 0
	for i, _ := range dataSlice {
		keys = append(keys, i)
	}
	sort.Ints(keys) // sort data index using a temp array
	for i, _ := range keys {
		max = i
	}
	// iterate through the index  array and use to print data
	if keys != nil {
		for key := keys[0]; key <= keys[max]; key = key + 4 {
			if (key-keys[0])%32 == 0 {
				fmt.Fprintf(f, "\n%d:\t", key)
				for i := 0; i < 32; i = i + 4 {
					if dataSlice[key+i] != 0 {
						fmt.Fprintf(f, "%d\t", dataSlice[key+i])
					} else {
						dataSlice[key+i] = 0
						fmt.Fprintf(f, "%d\t", dataSlice[key+i])
					}
				}
			}

		}
	}

	fmt.Fprintf(f, "\n")

}

func printPipeline(sim []Instruction, f *os.File, count int) {
	//f, fileErr := os.Create(file)
	//if fileErr != nil {
	//	fmt.Println(fileErr)
	//}

	fmt.Fprintln(f, "--------------------")
	fmt.Fprintf(f, "Cycle: %d\n", count)

	fmt.Fprintln(f, "\nPre-Issue Buffer:")
	fmt.Fprintf(f, "\tEntry 0:\t%s\n", fetchStr(sim, PreIssueBuff[0]))
	fmt.Fprintf(f, "\tEntry 1:\t%s\n", fetchStr(sim, PreIssueBuff[1]))
	fmt.Fprintf(f, "\tEntry 2:\t%s\n", fetchStr(sim, PreIssueBuff[2]))
	fmt.Fprintf(f, "\tEntry 3:\t%s", fetchStr(sim, PreIssueBuff[3]))

	fmt.Fprintln(f, "\nPre-ALU Queue:")
	fmt.Fprintf(f, "\tEntry 0:\t%s\n", fetchStr(sim, PreALUBuff[0]))
	fmt.Fprintf(f, "\tEntry 1:\t%s", fetchStr(sim, PreALUBuff[1]))

	fmt.Fprintln(f, "\nPost-ALU Queue:")
	fmt.Fprintf(f, "\tEntry 0:\t%s", fetchStr(sim, postALUBuff[1]))

	fmt.Fprintln(f, "\nPre-MEM Queue:")
	fmt.Fprintf(f, "\tEntry 0:\t%s\n", fetchStr(sim, PreMemBuff[0]))
	fmt.Fprintf(f, "\tEntry 1:\t%s", fetchStr(sim, PreMemBuff[1]))

	fmt.Fprintln(f, "\nPost-MEM Queue:")
	fmt.Fprintf(f, "\tEntry 0:\t%s\n", fetchStr(sim, postMemBuff[1]))

	// print current register
	fmt.Fprint(f, "\nRegisters:\n")
	fmt.Fprintf(f, "r00:\t%s", mapToString(registerMap, 8))
	fmt.Fprintf(f, "\nr08:\t%s", mapToString(registerMap, 16))
	fmt.Fprintf(f, "\nr16:\t%s", mapToString(registerMap, 24))
	fmt.Fprintf(f, "\nr24:\t%s\n", mapToString(registerMap, 32))

	fmt.Fprintln(f, "\nCache")
	fmt.Fprintf(f, "Set 0: LRU=%d\n", LRUbits[0])
	fmt.Fprintf(f, "\tEntry 0:%s\n", cacheStr(0, 0))
	fmt.Fprintf(f, "\tEntry 1:%s\n", cacheStr(0, 1))
	fmt.Fprintf(f, "Set 1: LRU=%d\n", LRUbits[1])
	fmt.Fprintf(f, "\tEntry 0:%s\n", cacheStr(1, 0))
	fmt.Fprintf(f, "\tEntry 1:%s\n", cacheStr(1, 1))
	fmt.Fprintf(f, "Set 2: LRU=%d\n", LRUbits[2])
	fmt.Fprintf(f, "\tEntry 0:%s\n", cacheStr(2, 0))
	fmt.Fprintf(f, "\tEntry 1:%s\n", cacheStr(2, 1))
	fmt.Fprintf(f, "Set 3: LRU=%d\n", LRUbits[3])
	fmt.Fprintf(f, "\tEntry 0:%s\n", cacheStr(3, 0))
	fmt.Fprintf(f, "\tEntry 1:%s\n", cacheStr(3, 1))

	// print data
	fmt.Fprintf(f, "\nData:")
	var keys []int
	max := 0
	for i, _ := range dataSlice {
		keys = append(keys, i)
	}
	sort.Ints(keys) // sort data index using a temp array
	for i, _ := range keys {
		max = i
	}
	// iterate through the index  array and use to print data
	if keys != nil {
		for key := keys[0]; key <= keys[max]; key = key + 4 {
			if (key-keys[0])%32 == 0 {
				fmt.Fprintf(f, "\n%d:\t", key)
				for i := 0; i < 32; i = i + 4 {
					if dataSlice[key+i] != 0 {
						fmt.Fprintf(f, "%d\t", dataSlice[key+i])
					} else {
						dataSlice[key+i] = 0
						fmt.Fprintf(f, "%d\t", dataSlice[key+i])
					}
				}
			}

		}
	}

	fmt.Fprintf(f, "\n")
}
