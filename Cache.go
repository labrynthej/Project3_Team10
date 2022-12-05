package main

import (
	"fmt"
)

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
var JustMissedList []int // (later)

var tagMask = 134217727 << 5
var setMask = 24

func writeToCache(count int) {
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

	if LRUbits[setNum] == 0 {
		dataWord = 0
	}
	if LRUbits[setNum] == 1 {
		dataWord = 1
	}
	LRUbits[setNum] = LRUbits[setNum] ^ 1
	CacheSets[setNum][dataWord].valid = 1
	CacheSets[setNum][dataWord].tag = tag
	CacheSets[setNum][dataWord].word1 = address1
	CacheSets[setNum][dataWord].word2 = address2

}

func readFromCache(count int) (int, int) {

	setNum := (count & setMask) >> 3
	//tag := (count & tagMask) >> 5
	index := 0
	//tag := (address1 & tagMask) >> 5

	if CacheSets[setNum][LRUbits[setNum]].word1 != count {
		writeToCache(count)
		//return -1, -1
	}
	// writeToCache(count)
	index = (count - 96) / 4

	//setNum := (count & setMask) >> 3
	//tag := (count & tagMask) >> 5

	// lineValue1, _ := strconv.ParseUint(memoryMap[count], 2, 32)
	// lineValue2, _ := strconv.ParseUint(memoryMap[count], 2, 32)

	return index, index + 1
}

func cacheStr(idx1 int, idx2 int) string {
	address1 := CacheSets[idx1][idx2].word1
	address2 := CacheSets[idx1][idx2].word2

	s := fmt.Sprintf("[(%d, %d, %d)<%s,%s>]", CacheSets[idx1][idx2].valid, CacheSets[idx1][idx2].dirty,
		CacheSets[idx1][idx2].tag, memoryMap[address1], memoryMap[address2])

	return s
}
