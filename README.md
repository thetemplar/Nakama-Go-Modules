## Get Dependencies
```
go get -u github.com/bradfitz/slice
go get -u github.com/gofrs/uuid
go get -u github.com/golang/protobuf/proto
```

## Install Protobuf
https://askubuntu.com/questions/1072683/how-can-i-install-protoc-on-ubuntu-16-04
###### Prerequesites
$ `sudo apt-get install autoconf automake libtool curl make g++ unzip`
###### Installation
1) https://github.com/protocolbuffers/protobuf/releases/tag/v3.6.1
2) Extract the contents and change in the directory
3) `./configure`
4) `make`
5) `make check`
6) `sudo make install`
7) `sudo ldconfig # refresh shared library cache.`
###### Check if it works
$ `protoc --version`


## Install gofast 
See: https://github.com/gogo/protobuf
```
go get github.com/gogo/protobuf/protoc-gen-gofast	
export PATH=$PATH:$GOPATH/bin
```

## make and run
`./make` (its a file, not the "real" make)
