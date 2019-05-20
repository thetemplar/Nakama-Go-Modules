#export GOROOT=/usr/local/go
#export GOPATH=$HOME/go                                                                                                                                                                                                         
#export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
rm modules/plugin.so &> /dev/null
protoc --gofast_out=. --csharp_out=. matchstate.proto 
#chmod 777 *.go
go build --buildmode=plugin -o ~/go/src/modules/plugin.so &&  ~/go/src/github.com/heroiclabs/nakama/nakama --runtime.path ~/go/src/Nakama-Go-Modules/modules  
#--logger.level "debug"           
