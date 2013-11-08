setenv NPGO_DIR /global/homes/n/npadmana/myWork/gocode/src/github.com/npadmana/npgo
module load fftw

#----- We should try to ensure that everything below is automatic
setenv PKG_CONFIG_PATH ${NPGO_DIR}/local/lib/pkgconfig:${MPICH_DIR}/lib/pkgconfig:${NPGO_DIR}/pkgconfigs:${NPGO_DIR}/pkgconfigs/edison
setenv LD_LIBRARY_PATH ${NPGO_DIR}/local/lib:${NPGO_DIR}/clibs:${LD_LIBRARY_PATH}
setenv PETSC_DIR ${NPGO_DIR}/local
setenv C_INCLUDE_PATH ${NPGO_DIR}/clibs

