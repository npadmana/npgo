setenv NPGO_DIR /global/homes/n/npadmana/myWork/gocode/src/github.com/npadmana/npgo
setenv MPI_DIR $MPICH_DIR

#----- We should try to ensure that everything below is automatic
setenv PKG_CONFIG_PATH ${NPGO_DIR}/local/lib/pkgconfig:${MPI_DIR}/lib:${NPGO_DIR}/pkgconfigs
setenv LD_LIBRARY_PATH ${NPGO_DIR}/local/lib:${NPGO_DIR}/clibs:${LD_LIBRARY_PATH}
setenv PETSC_DIR ${NPGO_DIR}/local
setenv C_INCLUDE_PATH $PNPGO_DIR}/clibs

