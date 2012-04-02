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

	var message int
	if (rank == 0) {
		message = 60
		for i := 1; i < size; i++ {
			fmt.Printf("root sending message to %d\n", i)
			mpi.Send(&message, 1, mpi.INT, i, 10, mpi.COMM_WORLD)
		}
	} else {
		mpi.Recv(&message, 1, mpi.INT, 0, 10, mpi.COMM_WORLD, mpi.STATUS_IGNORE)
	}
	mpi.Barrier(mpi.COMM_WORLD)
}
