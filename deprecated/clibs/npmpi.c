#include <stdlib.h>
#include "mpi.h"

MPI_Op mpiop(int i) {
        MPI_Op retval;
        switch(i) {
        case 0 :
                retval = MPI_SUM;
                break;
        default :
                MPI_Abort(MPI_COMM_WORLD,1);
        }
        return retval;
}

MPI_Datatype mpitype(int i) {
        MPI_Datatype retval;
        switch(i) {
        case 0 :
                retval = MPI_LONG;
                break;
        default :
                MPI_Abort(MPI_COMM_WORLD,1);
        }
        return retval;
}
