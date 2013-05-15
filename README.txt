golang-syso
===========

$ export GOPATH=`pwd`
$ go install -v ...
# uses internal linker and complains of bad relocation
$ go install -v -ldflags '-linkmode=external' ...
# uses gcc to link, no complaints
$ ./bin/foo
len=0, cap=0
[]

# Instead of 
len=1024, cap=1024
[...]
