package random

import (
	"math/rand"
	"time"
)

type Random interface {
	Next() int
}

type random struct {
	ceil int
}

func New(max int) Random {
	rand.Seed(time.Now().UnixMicro())
	return &random{
		ceil: max,
	}
}

func (r *random) Next() int {
	return rand.Int() % r.ceil
}
