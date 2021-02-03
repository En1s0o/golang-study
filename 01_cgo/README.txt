MinGW64:
gcc -shared -fPIC -I. native.c -o native.dll

Linux:
gcc -shared -fPIC -I. native.c -o libnative.so
