export NPGO_DIR=/Users/npadmana/myWork/gocode/src/github.com/npadmana/npgo

#----- We should try to ensure that everything below is automatic
export PKG_CONFIG_PATH=$NPGO_DIR/local/lib/pkgconfig:/usr/local/openmpi/v1.6-gcc47/lib/pkgconfig:$NPGO_DIR/pkgconfigs:$NPGO_DIR/pkgconfigs/laptop
export PETSC_DIR=$NPGO_DIR/local
export C_INCLUDE_PATH=$NPGO_DIR/clibs:

