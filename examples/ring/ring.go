package main

import (
	"mpi"
	"fmt"
)

func main() {
	defer mpi.Finalize()
	var message int
	var tag int = 201

	rank := mpi.Comm_rank(mpi.COMM_WORLD)
	size := mpi.Comm_size(mpi.COMM_WORLD)

	/* Calculate the rank of the next process in the ring.  Use the
	 modulus operator so that the last process "wraps around" to
	 rank zero. */

	next := (rank + 1) % size
	prev := (rank + size - 1) % size

	/* If we are the "master" process (i.e., MPI_COMM_WORLD rank 0),
	 put the number of times to go around the ring in the
	 message. */

	if (0 == rank) {
		message = 10
		fmt.Printf("Process 0 sending %d to %d, tag %d (%d processes in ring)\n", 
			message, next, tag, size)
		mpi.Send(&message, 1, mpi.INT, next, tag, mpi.COMM_WORLD)
		fmt.Printf("Process 0 sent to %d\n", next)
	}

	/* Pass the message around the ring.  The exit mechanism works as
	 follows: the message (a positive integer) is passed around the
	 ring.  Each time it passes rank 0, it is decremented.  When
	 each processes receives a message containing a 0 value, it
	 passes the message on to the next process and then quits.  By
	 passing the 0 message first, every process gets the 0 message
	 and can quit normally. */

	for true {
		mpi.Recv(&message, 1, mpi.INT, prev, tag, mpi.COMM_WORLD, mpi.STATUS_IGNORE)

		if (0 == rank) {
		        message--
		        fmt.Printf("Process 0 decremented value: %d\n", message)
		}

		mpi.Send(&message, 1, mpi.INT, next, tag, mpi.COMM_WORLD)
		if (0 == message) {
			fmt.Printf("Process %d exiting\n", rank)
			break
		}
	}

	/* The last process does one extra send to process 0, which needs
	   to be received before the program can exit */

	if (0 == rank) {
		mpi.Recv(&message, 1, mpi.INT, prev, tag, mpi.COMM_WORLD, mpi.STATUS_IGNORE)
	}
}
