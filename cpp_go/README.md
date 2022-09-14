add.cpp is an utility of CPP is called From Go

Steps to create a libadd.so file from add.cpp

gcc -c -o add.o add.cpp

gcc -shared -o libadd.so add.o

adder.c  - c code sample to use so library in c and call the add utility

In Go code, we are using cgo to call add utility of libadd.so
