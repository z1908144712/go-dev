package main

import (
	"day7/example1/balance"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// var insts []*balance.Instance
	insts := make([]*balance.Instance, 0)
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}
	balanceName := "hash"
	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
