setenv NPGO_DIR /global/homes/n/npadmana/myWork/gocode/src/github.com/npadmana/npgo

#----- We should try to ensure that everything below is automatic
setenv PKG_CONFIG_PATH ${NPGO_DIR}/local/lib/pkgconfig:${PKG_CONFIG_PATH}:${NPGO_DIR}/pkgconfigs
setenv LD_LIBRARY_PATH ${NPGO_DIR}/local/lib:${NPGO_DIR}/clibs:${LD_LIBRARY_PATH}
export PETSC_DIR=$NPGO_DIR/local
export C_INCLUDE_PATH=$NPGO_DIR/clibs:$C_INCLUDE_PATH

