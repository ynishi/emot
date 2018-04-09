.PHONY: compile 

compile:
	protoc -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. *.proto
	protoc -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. *.proto
	goimports -w server/*.go 
	cd server && go build server.go
	cd server && go build proxy.go
