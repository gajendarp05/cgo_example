1. C have completely different heap which we can use to allocate memory.
2. Golang GC will not know about C memory and hence cannot Garbage collect it.
3. Programmer have a responsibility to manually free the pointer after use using C.free()
