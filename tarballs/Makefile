clean :
	- rm -rf build


gsl :
	mkdir -p build
	cd build; \
	tar xvfz ../gsl-1.16.tar.gz; \
	cd gsl-1.16; \
	./configure --prefix=$(NPGO_DIR)/local; \
	make; \
	make install 



petsc :
	mkdir -p build
	cd build; \
	tar xvfz ../petsc-lite-3.4.3.tar.gz; \
	cd petsc-3.4.3; \
	unset PETSC_DIR; \
	./configure --prefix=$(NPGO_DIR)/local \
	   --with-shared-libraries=1 --with-fortran=0 \
	   --with-64-bit-indices=1 --with-cc=mpicc --with-debugging=0 --with-cxx=mpicxx \
	   --download-f2cblaslapack; \
	make PETSC_DIR=$(NPGO_DIR)/tarballs/build/petsc-3.4.3 PETSC_ARCH=arch-darwin-c-opt all; \
	make PETSC_DIR=$(NPGO_DIR)/tarballs/build/petsc-3.4.3 PETSC_ARCH=arch-darwin-c-opt install


fftw :
	mkdir -p build
	cd build; \
	tar xvfz ../fftw-3.3.3.tar.gz; \
	cd fftw-3.3.3; \
	./configure --prefix=$(NPGO_DIR)/local \
		--enable-mpi --enable-shared; \
	make; \
	make install