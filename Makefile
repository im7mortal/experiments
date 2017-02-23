INSTALL_DIR = /usr/local/

all:
	go build -buildmode=c-shared -o libcglog.so cglog.go

clean:
	rm libcglog.so
	rm libcglog.h

install: $(TARGET)
	cp libcglog.so ${INSTALL_DIR}lib/libcglog.so
	cp libcglog.h ${INSTALL_DIR}include/libcglog.h
