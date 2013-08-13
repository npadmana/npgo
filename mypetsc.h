#include "petsc.h"

int sizePetscReal();
int sizePetscInt();
void petscAbort();
char* petscNullChar();

PetscErrorCode mypetscPrintf(char* s);
PetscErrorCode mypetscSyncPrintf(char* s);