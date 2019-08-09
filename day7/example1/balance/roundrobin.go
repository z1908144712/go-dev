package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No Instance")
		return
	}
	lens := len(insts)
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
