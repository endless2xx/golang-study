package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {

	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}
