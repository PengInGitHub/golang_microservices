# golang_microservices
Exercise of Microservices in Golang by Ewan Valentine<br/> 
Make use of protobuf and gRPC as transport protocol<br/>  
Microservice deals with separating complexity<br/> 
Golang: light, fast, support concurrency<br/> 
go-micro: the microservice framework available for Go<br/> 
communication is very important in microservice since microservices are separated<br/> 
gRPC:light-weight binary based RPC communication protocol made by Google<br/> 
protobuf:the interchange DSL of gRPC, used to define interface to service<br/> 
Protocol Buffers v3:the protoc compiler that is used to generate gRPC service code<br/> 

To note, pay attention to make file in tutorial one: the path in Makefile may be wrong, must be adapted to your machine.<br/> 
.proto is a protobuf definition file that defines a gRPC service<br/> 
GRPC is a high performance, open-source, universial RPC framework<br/> 
Compiling the .proto with protocol tool with the grpc plugin will marshal/unmarshal the msg btw Go code and buffer protocol buffer binary msg.

