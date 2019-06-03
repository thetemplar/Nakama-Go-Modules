#export GOROOT=/usr/local/go
#export GOPATH=$HOME/go                                                                                                                                                                                                         
#export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
rm modules/plugin.so &> /dev/null
cd protobuf
protoc --gofast_out=../ --csharp_out=. matchstate.proto
protoc --gofast_out=../ --csharp_out=. character.proto
cd ..
cp GameDB/GameDB* .
go build --buildmode=plugin -o ~/go/src/modules/plugin.so && rm GameDB*.go && ~/go/src/github.com/heroiclabs/nakama/nakama --runtime.path ~/go/src/modules
#--logger.level "debug"      
rm GameDB*.go  &>/dev/null     

