// Copyright 2012 Abhishek Kulkarni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !mpich

package mpi

/*
#cgo LDFLAGS: -L/Users/adkulkar/opt/lib/ -lmpi
#cgo CFLAGS: -I/Users/adkulkar/opt/include/

#include <mpi.h>
*/
import "C"

var COMM_WORLD MPI_Comm = &C.ompi_mpi_comm_world
var COMM_SELF MPI_Comm = &C.ompi_mpi_comm_self

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
