// Copyright 2012 Abhishek Kulkarni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mpi

/*
#cgo LDFLAGS: -L/Users/adkulkar/opt/lib/ -lmpi
#cgo CFLAGS: -I/Users/adkulkar/opt/include/

#include <stdlib.h>
#include <mpi.h>

*/
import "C"

import (
	"os"
	"unsafe"
)

// Communicator
type MPI_Comm *[0]uint8

// Datatypes
type MPI_Datatype *[0]uint8

type MPI_Status *C.MPI_Status

var STATUS_IGNORE MPI_Status = nil
var STATUSES_IGNORE MPI_Status = nil

func Init() {
	var argc int
	C.MPI_Init( (*C.int)(unsafe.Pointer(&argc)), (***C.char)(unsafe.Pointer(&os.Args)))
}

func Comm_rank(comm MPI_Comm) int {
	var rank int
	C.MPI_Comm_rank(comm, (*C.int)(unsafe.Pointer(&rank)) )
	return rank
}

func Comm_size(comm MPI_Comm) int {
	var size int
	C.MPI_Comm_size(comm, (*C.int)(unsafe.Pointer(&size)) )
	return size
}

func Send(buf interface{}, count int, datatype MPI_Datatype, dest int,
	tag int, comm MPI_Comm) int {
	switch ptr := buf.(type) {
	case *int:
		ret := C.MPI_Send(unsafe.Pointer(ptr), C.int(count), datatype,
			C.int(dest), C.int(tag), comm)
		return int(ret)
	}
	return 0
}

func Recv(buf interface{}, count int, datatype MPI_Datatype, source int, 
	tag int, comm MPI_Comm, status MPI_Status) int {
	switch ptr := buf.(type) {
	case *int:
		ret := C.MPI_Recv(unsafe.Pointer(ptr), C.int(count), datatype,
			C.int(source), C.int(tag), comm, status)
		return int(ret)
	}
	return 0
}

func Barrier(comm MPI_Comm) int {
	ret := C.MPI_Barrier(comm)
	return int(ret)
}

func Finalize() {
	C.MPI_Finalize()
}
