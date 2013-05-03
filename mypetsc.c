#include "petsc.h"

int sizePetscReal() {
	return sizeof(PetscScalar);
}

int sizePetscInt() {
	return sizeof(PetscInt);
}

void petscAbort() {
	MPI_Abort(PETSC_COMM_WORLD, -1);
}

char* petscNullChar() {
	return (char*) PETSC_NULL;
};

void petscRankSize(int *rank, int *size) {
	MPI_Comm_rank(PETSC_COMM_WORLD, rank);
	MPI_Comm_size(PETSC_COMM_WORLD, size);
}

PetscErrorCode mypetscPrintf(char* s) {
	return PetscPrintf(PETSC_COMM_WORLD, s);
}
PetscErrorCode mypetscSyncPrintf(char* s) {
	return PetscSynchronizedPrintf(PETSC_COMM_WORLD, s);
}