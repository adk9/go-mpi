// Copyright 2012 Abhishek Kulkarni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"mpi"
	"fmt"
)

func main() {
	mpi.Init()
	defer mpi.Finalize()

	rank := mpi.Comm_rank(mpi.COMM_WORLD)
	size := mpi.Comm_size(mpi.COMM_WORLD)
	fmt.Printf("Hello, world, I am %d of %d\n", rank, size)
	mpi.Barrier(mpi.COMM_WORLD)
}