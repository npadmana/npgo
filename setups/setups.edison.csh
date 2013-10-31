setenv NPGO_DIR /global/homes/n/npadmana/myWork/gocode/src/github.com/npadmana/npgo

#----- We should try to ensure that everything below is automatic
setenv PKG_CONFIG_PATH ${NPGO_DIR}/local/lib/pkgconfig:${PKG_CONFIG_PATH}
setenv LD_LIBRARY_PATH ${NPGO_DIR}/local/lib:${LD_LIBRARY_PATH}

