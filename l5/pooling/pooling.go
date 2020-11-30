package pooling

import (
	"Middleware-IF711-2020-3/l5/hashing"
)

type Pool struct {
	Instances []hashing.Hash
}

func AllocatePool(size int) Pool {
	returnPool := Pool{}
	returnPool.Instances = make([]hashing.Hash, size)
	return returnPool
}
