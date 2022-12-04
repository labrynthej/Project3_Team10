package main

type Block struct {
	valid int
	dirty int
	tag   int
	word1 int
	word2 int
}

var Set = [2]Block{
	Block{
		valid: 0,
		dirty: 0,
		tag:   0,
		word1: 0,
		word2: 0},
	Block{
		valid: 1,
		dirty: 1,
		tag:   1,
		word1: 2,
		word2: 1,
	},
}
var CacheSets [4][2]Block
var LRUbits = [4]int{0, 0, 0, 0}

var tagMask = 134217727 << 5
var setMask = 24

func writeToCache(instArray []Instruction, count int) {
	// index := (count - 96) / 4
	var dataWord int
	var address1, address2 int

	if count%8 == 0 { //  then its aligned correct
		dataWord = 0 // block 0 was the address
		address1 = count
		address2 = count + 4
	}
	if count%8 != 0 { // not aligned correctly
		dataWord = 1 // block 1 was the address
		address1 = count - 4
		address2 = count
	}

	setNum := (address1 & setMask) >> 3
	tag := (address1 & tagMask) >> 5
	LRUbits[setNum] = dataWord ^ 1

	CacheSets[setNum][dataWord].valid = 0
	CacheSets[setNum][dataWord].tag = tag
	CacheSets[setNum][dataWord].word1 = address1
	CacheSets[setNum][dataWord].word2 = address2

}
