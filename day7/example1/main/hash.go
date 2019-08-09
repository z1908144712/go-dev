package main

import (
	"day7/example1/balance"
	"errors"
	"hash/crc32"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance) (inst *balance.Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("lens Error")
		return
	}
	defKey := "hash"
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashValue := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashValue) % lens
	inst = insts[index]
	return
}
