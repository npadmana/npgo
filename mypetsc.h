#include "petsc.h"

int sizePetscReal();
int sizePetscInt();
void petscAbort();
char* petscNullChar();

void petscRankSize(int *rank, int *size);

PetscErrorCode mypetscPrintf(char* s);
PetscErrorCode mypetscSyncPrintf(char* s);