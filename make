export GOROOT=/usr/local/go
export GOPATH=$HOME/go                                                                                                                                                                                                         
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
gofmt -w init.go
rm modules/plugin.so &> /dev/null
cd protobuf
protoc --gofast_out=../ --csharp_out=. matchstate.proto
protoc --gofast_out=../ --csharp_out=. character.proto
cd ..
go build -trimpath --buildmode=plugin -o ~/go/src/modules/plugin.so && ~/Nakama-Repo/nakama/nakama --runtime.path ~/go/src/modules
#--logger.level "debug"        
