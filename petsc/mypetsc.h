#include "petsc.h"

int sizePetscReal();
int sizePetscInt();
void petscAbort();
char* petscNullChar();
Vec petscNullVec();


PetscErrorCode mypetscPrintf(char* s);
PetscErrorCode mypetscSyncPrintf(char* s);
PetscErrorCode setTypeMPI(Vec v);