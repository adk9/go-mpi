package main

import (
	"mpi"
	"fmt"
)

func main() {
	defer mpi.Finalize()
	rank := mpi.Comm_rank(mpi.COMM_WORLD)
	size := mpi.Comm_size(mpi.COMM_WORLD)
	fmt.Printf("Hello, world, I am %d of %d\n", rank, size)
	mpi.Barrier(mpi.COMM_WORLD)
}