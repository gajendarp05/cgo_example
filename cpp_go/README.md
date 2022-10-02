__add.cpp__ is an utility of CPP is called From Go

**Steps to create a libadd.so file from add.cpp**

1. gcc -c -o add.o add.cpp

2. gcc -shared -o libadd.so add.o

__adder.c__  - C code sample to use so library in c and call the add utility

In Go code, we are using ***cgo*** to call add utility of libadd.so
