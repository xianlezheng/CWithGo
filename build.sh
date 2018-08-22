go build -buildmode=c-archive lib.go
gcc -pthread test.c lib.a -o test
./test 123456