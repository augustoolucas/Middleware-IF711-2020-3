package pooling

import (
	"Middleware-IF711-2020-3/l5/hashing"
	"shared"
)

type Pool struct {
	Instances []hashing.Hash
}

func (pool Pool) AllocatePool(size int) Pool {
	pool.Instances = make([]hashing.Hash, size)
	for i := 0; i < shared.POOL_SIZE; i++ {
		pool.Instances[i].Available = true
	}
	return pool
}
