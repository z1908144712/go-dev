package balance

import (
	"errors"
)

type BalanceMgr struct {
	allBalances map[string]Balancer
}

var mgr = BalanceMgr{
	allBalances: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalancer(name string, b Balancer) {
	p.allBalances[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalances[name]
	if !ok {
		err = errors.New("No Balancer Found")
		return
	}
	inst, err = balancer.DoBalance(insts)
	return
}
