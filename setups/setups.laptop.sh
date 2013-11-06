export NPGO_DIR=/Users/npadmana/myWork/gocode/src/github.com/npadmana/npgo

#----- We should try to ensure that everything below is automatic
export PKG_CONFIG_PATH=$NPGO_DIR/local/lib/pkgconfig:$PKG_CONFIG_PATH:$NPGO_DIR/pkgconfigs
export PETSC_DIR=$NPGO_DIR/local
export C_INCLUDE_PATH=$NPGO_DIR/clibs:$C_INCLUDE_PATH
export DYLD_LIBRARY_PATH=$NPGO_DIR/clibs:$DYLD_LIBRARY_PATH
export LD_LIBRARY_PATH=$NPGO_DIR/clibs:$LD_LIBRARY_PATH

