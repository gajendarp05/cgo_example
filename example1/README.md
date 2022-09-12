C.CString coverts Go to C strings and this conversions make a copy of data in C memory
Since memory allocations in C are not known to Go's memory manager, always free the C memory manually when you are done.
The returned pointer to a char array by C is converted to unsafe.Pointer
Memory is freed by calling C.free

References:
https://go.dev/blog/cgo
https://pkg.go.dev/unsafe#Pointer