#include <iostream>

extern "C" int Add(int x, int y) {
    printf("Adding...  %d, %d\n", x, y);
    return x + y;
}