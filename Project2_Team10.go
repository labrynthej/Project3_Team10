package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Instruction struct {
	typeOfInstruction string // instruction type "R", "I", etc..
	rawInstruction    string // raw data (we need to run this through a function that figures out the OPcode)
	lineValue         uint64 // linevalue = rawinstruction converted to uint64 so I could use mask and shift on it.
	field             uint64
	opcode            uint64 // once we know this we can figure out everything else
	op                string // what is it? ADD, SUB, LSL, etc...
	rd                uint8
	rn                uint8
	rm                uint8
	im                uint8
	rt                uint8
	address           uint8
	offset            int32
	conditional       uint8
	shamt             uint8
	op2               uint8
	cycle             int
	programCnt        int // program counter
	data              int32
	destReg           int // this is rd
	src1Reg           int //this is rn
	src2Reg           int // this is rm
	arg1Str           string
	arg2Str           string
	arg3Str           string
	isMemOp           bool
}

// global data slice
var dataSlice = make(map[int]int)

// global register map
var registerMap = make(map[uint8]int)

// global memory map
var memoryMap = make(map[int]string)

func main() {
	//flag.String gets pointers to command line arguments
	cmdInFile := flag.String("i", "addtest1_bin.txt", "-i [input file path/name]")
	cmdOutFile := flag.String("o", "team10_out.txt", "-o [output file path/name]")
	flag.Parse() //flag.parse just makes things work

	inFile, _ := os.Open(*cmdInFile) //*cmdInFile because we need to dereference it

	//create a new array of instructions based on the data read from the inFile
	var instructionsArray []Instruction = readFile(inFile)
	initializeInstructions(instructionsArray) //initialize the instructions

	printResults(instructionsArray, *cmdOutFile+"_dis.txt")

	// begin simulation
	//simInstructions(instructionsArray, *cmdOutFile+"_sim.txt")

	// begin pipeline simulation
	controlUnit(instructionsArray)
	//printPipeline(instructionsArray, *cmdOutFile+"_test.txt")

	fmt.Println("infile:", *cmdInFile)
	fmt.Println("outfile: ", *cmdOutFile+"_dis.txt")
	//fmt.Println("simulation outfile: ", *cmdOutFile+"_sim.txt")

}

// reads the file and loads each line into the rawInstruction part of the Instruction
func readFile(fileBeingRead io.Reader) (inputParsed []Instruction) {
	index := 0

	// ^^ begins reading from the fileBeingRead (io.Reader is like an iostream) ...
	// ... and prepares an array of Instructions that will be returned later
	scanner := bufio.NewScanner(fileBeingRead) // "scanner" is a bufio.NewScanner of the file being read. ...

	// ...By default, a "Scan" terminates at new line.
	for scanner.Scan() { // for each Scan (each line to read) in "scanner"
		newInstruction := Instruction{rawInstruction: scanner.Text()} // creates a new Instruction and assigns...
		// ...scanner.Text() to rawInstruction (scanner.Text() is the text of one line of the input file)
		inputParsed = append(inputParsed, newInstruction) // add the newInstruction (containing the raw data) to the...
		// ...array of instructions

		// set the current memory value then increment by 4
		inputParsed[index].programCnt = 96 + (4 * index)
		memoryMap[inputParsed[index].programCnt] = inputParsed[index].rawInstruction
		index++

	}
	fmt.Println(scanner.Text())
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return
}

// intialize all values of the struct array based on the line value
// mask bits using with bitwise-AND value to hex value of mask
// ie. "lineValue & 0xFF00" with a 32-bit mask 0xFF00 == 00000000000000001111111100000000
// shift bits using >> operation
func initializeInstructions(instArray []Instruction) {
	for i := 0; i < len(instArray); i++ {
		//the below converts 32 characters from a base 2 string to base 10 uint64

		lineValue, _ := strconv.ParseUint(instArray[i].rawInstruction, 2, 32)

		if lineValue > 335544320 || lineValue == 0 {
			// assign lineValue and 11 bit opcode for setting the instruction
			instArray[i].lineValue = lineValue
			instArray[i].opcode = lineValue >> 21
			setInstructionType(instArray, i)

			// set values for instruction type "R" | opcode | Rm | Shamt | Rn | Rd |
			if instArray[i].typeOfInstruction == "R" {
				instArray[i].rn = uint8((lineValue & 0x3E0) >> 5)
				instArray[i].rm = uint8((lineValue & 0x1F0000) >> 16)
				instArray[i].rd = uint8(lineValue & 0x1F)
				instArray[i].shamt = uint8((lineValue & 0xFC00) >> 11)
			}

			// set values for instruction type "D" | opcode | address | op2 | Rn | Rt |
			if instArray[i].typeOfInstruction == "D" {
				instArray[i].rn = uint8((lineValue & 0x3E0) >> 5)
				instArray[i].address = uint8((lineValue & 0x1FF000) >> 12)
				instArray[i].op2 = uint8((lineValue & 0xC00) >> 10)
				instArray[i].rt = uint8(lineValue & 0x1F)
			}

			// set values for instruction type "I" | opcode | immediate | Rn | Rd |
			if instArray[i].typeOfInstruction == "I" {
				instArray[i].opcode = lineValue >> 22
				instArray[i].rn = uint8((lineValue & 0x3E0) >> 5)
				instArray[i].im = uint8(signedVariable(lineValue&0x3FFC00>>10, 12))
				instArray[i].rd = uint8(lineValue & 0x1F)
			}

			// set values for instruction type "B" | opcode | offset |
			if instArray[i].typeOfInstruction == "B" {
				instArray[i].opcode = lineValue >> 26
				instArray[i].offset = signedVariable(lineValue&0x3FFFFFF, 26)
			}

			// set values for instruction type "CB" (conditional B) | opcode | offset |
			if instArray[i].typeOfInstruction == "CB" {
				instArray[i].opcode = lineValue >> 24
				instArray[i].offset = signedVariable(lineValue&0xFFFFE0>>5, 19)
				instArray[i].conditional = uint8(lineValue & 0x1F)
			}

			// set values for instruction type "IM" | opcode | shift | field | Rd |
			if instArray[i].typeOfInstruction == "IM" {
				instArray[i].opcode = lineValue >> 23
				instArray[i].shamt = uint8(lineValue & 0x600000 >> 21)
				instArray[i].field = lineValue & 0x1FFFE0 >> 5
				instArray[i].rd = uint8(lineValue & 0x1F)
			}

			if instArray[i].op == "BREAK" {
				break
			}
		}

	}
}

// check for signed variable to convert to negation using two's complement
func signedVariable(value uint64, length int) int32 {
	//var mask uint64 = 1 << (length - 1) // set mask to get sign
	var temp = value >> (length - 1)

	if temp == 1 { //
		value = value | (0xFFFFFFFF << length)
		//value = ^value + 1
	}
	return int32(value)
}

// function that determines the type of instruction and what it is
func setInstructionType(instrArray []Instruction, i int) {
	var decimalOPC uint64 = instrArray[i].opcode
	switch true { //switch defines a base case to test against.
	case ((decimalOPC >= 160) && (decimalOPC <= 191)): //if case == switch, do stuff in that one and ignore other cases
		instrArray[i].op = "B"
		instrArray[i].typeOfInstruction = "B"
	case (decimalOPC == 1104):
		instrArray[i].op = "AND"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC == 1112):
		instrArray[i].op = "ADD"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC >= 1160 && decimalOPC <= 1161):
		instrArray[i].op = "ADDI"
		instrArray[i].typeOfInstruction = "I"
	case (decimalOPC == 1360):
		instrArray[i].op = "ORR"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC >= 1440 && decimalOPC <= 1447):
		instrArray[i].op = "CBZ"
		instrArray[i].typeOfInstruction = "CB"
	case (decimalOPC >= 1448 && decimalOPC <= 1455):
		instrArray[i].op = "CBNZ"
		instrArray[i].typeOfInstruction = "CB"
	case (decimalOPC == 1624):
		instrArray[i].op = "SUB"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC >= 1672 && decimalOPC <= 1673):
		instrArray[i].op = "SUBI"
		instrArray[i].typeOfInstruction = "I"
	case (decimalOPC >= 1684 && decimalOPC <= 1687):
		instrArray[i].op = "MOVZ"
		instrArray[i].typeOfInstruction = "IM"
	case (decimalOPC >= 1940 && decimalOPC <= 1943):
		instrArray[i].op = "MOVK"
		instrArray[i].typeOfInstruction = "IM"
	case (decimalOPC == 1690):
		instrArray[i].op = "LSR"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC == 1691):
		instrArray[i].op = "LSL"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC == 1984):
		instrArray[i].op = "STUR"
		instrArray[i].typeOfInstruction = "D"
	case (decimalOPC == 1986):
		instrArray[i].op = "LDUR"
		instrArray[i].typeOfInstruction = "D"
	case (decimalOPC == 1692):
		instrArray[i].op = "ASR"
		instrArray[i].typeOfInstruction = "R"
	case (decimalOPC == 1872):
		instrArray[i].op = "EOR"
		instrArray[i].typeOfInstruction = "R"
	case decimalOPC == 0:
		instrArray[i].op = "NOP"
		instrArray[i].typeOfInstruction = "N/A"
	case decimalOPC == 2038: //check for break
		instrArray[i].op = "BREAK"
		instrArray[i].typeOfInstruction = "BREAK"
	default:
		break
	}
}

// simulation functions
func simInstructions(instrArray []Instruction, fileName string) {
	// run function to decide outcome then assign based on cycle

	// initialize all keys of map and set to 0 (empty registers)
	for j := 0; j < 32; j++ {
		registerMap[uint8(j)] = 0
	}

	cycle := 0 // initiliaze cycle

	// create the file and keep open until the loop closes
	file, fileErr := os.Create(fileName)
	if fileErr != nil {
		fmt.Println(fileErr)
	}

	// as long as instruction is not break, loop through all cycles
	i := 0
	for instrArray[i].typeOfInstruction != "BREAK" {
		count := 1
		switch instrArray[i].op {
		// R format instructions
		case "SUB": // 	rd = rn - rm
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] - registerMap[instrArray[i].rm]
			break
		case "AND": // rd = rm & rn
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] & registerMap[instrArray[i].rm]
			break
		case "ADD": // rd = rm + rn
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] + registerMap[instrArray[i].rm]
			break
		case "ORR": // rd = rm | rn
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] | registerMap[instrArray[i].rm]
			break
		case "EOR": // rd = rm ^ rn
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] ^ registerMap[instrArray[i].rm]
			break
		case "LSR": // rn shifted shamt
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] >> registerMap[instrArray[i].shamt]
			break
		case "LSL": // rd = rn << shamt
			registerMap[instrArray[i].rd] = (registerMap[instrArray[i].rn]) << registerMap[instrArray[i].shamt]
			break
		case "ASR": // rd = rn >> shamt pad with sign bit
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] >> registerMap[instrArray[i].shamt]
			break

		// D format instructions
		case "LDUR":
			registerMap[instrArray[i].rt] = dataSlice[registerMap[instrArray[i].rn]+int(instrArray[i].address)*4]
			break
		case "STUR":
			dataSlice[registerMap[instrArray[i].rn]+int(instrArray[i].address)*4] = registerMap[instrArray[i].rt]
			break

		// I format instructions
		case "ADDI": // rd = rn + im
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] + int(instrArray[i].im)
			break
		case "SUBI": // rd = rn - im
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rn] - int(instrArray[i].im)
			break

		// B and CB format instructions
		case "B": // PC = PC +- (4 * offset)
			count = int(instrArray[i].offset)
			break
		case "CBZ": // if (conditional == 0) {PC = 4 * offset}
			if instrArray[i].conditional == 0 {
				count = int(instrArray[i].offset)
			}
			break
		case "CBNZ": // if (conditional == 1) {PC = 4 * offset}
			if instrArray[i].conditional != 0 {
				count = int(instrArray[i].offset)
			}
			break

		// IM format instructions
		case "MOVZ":
			registerMap[instrArray[i].rd] = 0
			registerMap[instrArray[i].rd] = int(instrArray[i].field<<(instrArray[i].shamt*16)) &
				(0xFFFFFFFF << (instrArray[i].shamt * 16))
			break
		case "MOVK":
			registerMap[instrArray[i].rd] = registerMap[instrArray[i].rd] +
				int(instrArray[i].field<<(instrArray[i].shamt*16))
			break
		case "NOP":
			break
		}

		cycle++                              // increment cycle
		instrArray[i].cycle = cycle          // assign cycle to struct
		printSimulation(instrArray[i], file) // print current struct simulation
		i = i + count                        // increment loop counter
	}
	if instrArray[i].typeOfInstruction == "BREAK" {
		cycle++
		instrArray[i].cycle = cycle
		printSimulation(instrArray[i], file)
	}
}

func instructionString(sim Instruction) string {
	switch sim.typeOfInstruction {
	case "R":
		switch sim.op {
		case "LSL", "LSR", "ASR":
			return fmt.Sprintf("%s\tR%d, R%d, #%d", sim.op, sim.rd, sim.rm, sim.shamt)
		default:
			return fmt.Sprintf("%s\tR%d, R%d, R%d", sim.op, sim.rd, sim.rm, sim.rn)
		}
	case "I":
		return fmt.Sprintf("%s\tR%d, R%d, #%d", sim.op, sim.rd, sim.rn, sim.im)
	case "D":
		return fmt.Sprintf("%s\tR%d, [R%d, #%d]", sim.op, sim.rt, sim.rn, sim.address)
	case "B":
		return fmt.Sprintf("%s\t #%d", sim.op, sim.offset)
	case "CB":
		return fmt.Sprintf("%s\tR%d, #%d", sim.op, sim.conditional, sim.offset)
	case "IM":
		return fmt.Sprintf("%s\tR%d, %d, LSL %d", sim.op, sim.rd, sim.field, sim.shamt*16)
	default:
		return fmt.Sprintf("%s\t", sim.op)

	}
}

func mapToString(arr map[uint8]int, highValue uint8) string {
	var str = ""
	var i uint8
	for i = highValue - 8; i < highValue; i++ {
		str = str + strconv.Itoa(arr[i]) + "\t"
	}
	return str
}
