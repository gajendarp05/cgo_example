#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

int main() {
    void *handle;
    handle = dlopen("./libadd.so", RTLD_LAZY);
    if(!handle) {
        fprintf(stderr, "%s\n", dlerror());
        exit(1);
    }

    char *error;
    void(*addd);
    addd = dlsym(handle, "Add");
    if ((error = dlerror()) != NULL)  {
        fprintf (stderr, "%s\n", error);
        exit(1);
    }
    int (*Add)(int, int);
    Add = addd;
    printf("%d\n", (*Add)(4, 5));
    dlclose(handle);
    return 0;
}