```
$ go build -o libgreeter.so -buildmode=c-shared greeter.go
$ gcc -o greeter.o greeter.c ./libgreeter.so
$ gcc greeter.c -o greeter -L. -lgreeter
$ chmod 755 greeter
$ LD_LIBRARY_PATH=`pwd`:$LD_LIBRARY_PATH ./greeter
```
