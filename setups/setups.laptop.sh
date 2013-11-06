export NPGO_DIR=/Users/npadmana/myWork/gocode/src/github.com/npadmana/npgo
export MPI_DIR=/usr/local/openmpi/v1.6-gcc47/

#----- We should try to ensure that everything below is automatic
export PKG_CONFIG_PATH=$NPGO_DIR/local/lib/pkgconfig:$MPI_DIR/lib:$NPGO_DIR/pkgconfigs
export PETSC_DIR=$NPGO_DIR/local
export C_INCLUDE_PATH=$NPGO_DIR/clibs:
export DYLD_LIBRARY_PATH=$NPGO_DIR/clibs:$DYLD_LIBRARY_PATH
export LD_LIBRARY_PATH=$NPGO_DIR/clibs:$LD_LIBRARY_PATH

