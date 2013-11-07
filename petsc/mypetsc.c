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
	return (char*) NULL;
};

Vec petscNullVec() {
	return (Vec) NULL;
}

PetscErrorCode mypetscPrintf(char* s) {
	return PetscPrintf(PETSC_COMM_WORLD, s);
}
PetscErrorCode mypetscSyncPrintf(char* s) {
	return PetscSynchronizedPrintf(PETSC_COMM_WORLD, s);
}

PetscErrorCode setTypeMPI(Vec v) {
	return VecSetType(v, VECMPI);
}