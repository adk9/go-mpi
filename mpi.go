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

var COMM_WORLD MPI_Comm = &C.ompi_mpi_comm_world
var COMM_SELF MPI_Comm = &C.ompi_mpi_comm_self

// Datatypes
type MPI_Datatype *[0]uint8

var DATATYPE_NULL MPI_Datatype = &C.ompi_mpi_datatype_null
var BYTE MPI_Datatype = &C.ompi_mpi_byte
var PACKED MPI_Datatype = &C.ompi_mpi_packed
var CHAR MPI_Datatype = &C.ompi_mpi_char
var SHORT MPI_Datatype = &C.ompi_mpi_short
var INT MPI_Datatype = &C.ompi_mpi_int
var LONG MPI_Datatype = &C.ompi_mpi_long
var FLOAT MPI_Datatype = &C.ompi_mpi_float
var DOUBLE MPI_Datatype = &C.ompi_mpi_double
var LONG_DOUBLE MPI_Datatype = &C.ompi_mpi_long_double
var UNSIGNED_CHAR MPI_Datatype = &C.ompi_mpi_unsigned_char
var SIGNED_CHAR MPI_Datatype = &C.ompi_mpi_signed_char
var UNSIGNED_SHORT MPI_Datatype = &C.ompi_mpi_unsigned_short
var UNSIGNED_LONG MPI_Datatype = &C.ompi_mpi_unsigned_long
var UNSIGNED MPI_Datatype = &C.ompi_mpi_unsigned
var FLOAT_INT MPI_Datatype = &C.ompi_mpi_float_int
var DOUBLE_INT MPI_Datatype = &C.ompi_mpi_double_int
var LONG_DOUBLE_INT MPI_Datatype = &C.ompi_mpi_longdbl_int
var LONG_INT MPI_Datatype = &C.ompi_mpi_long_int
var SHORT_INT MPI_Datatype = &C.ompi_mpi_short_int
var TWO_INT MPI_Datatype = &C.ompi_mpi_2int
var UB MPI_Datatype = &C.ompi_mpi_ub
var LB MPI_Datatype = &C.ompi_mpi_lb
var WCHAR MPI_Datatype = &C.ompi_mpi_wchar
var LONG_LONG_INT MPI_Datatype = &C.ompi_mpi_long_long_int
var LONG_LONG MPI_Datatype = &C.ompi_mpi_long_long_int
var UNSIGNED_LONG_LONG MPI_Datatype = &C.ompi_mpi_unsigned_long_long
var TWO_COMPLEX MPI_Datatype = &C.ompi_mpi_2cplex
var TWO_DOUBLE_COMPLEX MPI_Datatype = &C.ompi_mpi_2dblcplex

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
