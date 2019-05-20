export GOROOT=/usr/local/go
export GOPATH=$HOME/go                                                                                                                                                                                                         
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 

rm project.so &> /dev/null
protoc --go_out=. --csharp_out=. matchstate.proto 
chmod 777 *.go
go build --buildmode=plugin -o ~/go/src/modules/project.so &&  ~/go/src/github.com/heroiclabs/nakama/nakama --runtime.path ~/go/src/modules  
#--logger.level "debug"           
