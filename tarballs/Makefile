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