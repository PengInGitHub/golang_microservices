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
Compiling the .proto with protocol tool with the grpc plugin will marshal/unmarshal the msg btw Go code and buffer protocol buffer binary msg.<br/> 
cli.go: an interface takes a JSON consignment file and interact with gRPC service<br/> 
Protocol Buffers: protocol buffers is a method of serializing structured data. It is useful in developing programs to communicate with each other over a wire or for storing data.<br/> 
How Protocol Buffers works: In a proto definition file (.proto) developer defines the service and data structure that called message and compile it with protoc. This compilation generates code could be invoked by a sended or recipient of these data structures. For example, example.proto will produce example.pb.cc and example.pb.h, which will define c++ classes for each message and service that example.proto defines.<br/> 
What-to-do in proto definition file:first of all, define the service, this should contain the mehtods you wish to expose to other services. Then you define the messages, which are effectively the data structure.<br/> 
In proto, message is handled by protobuf and service is handled by gRPC plugin which compiles code to interact with service.<br/> 
Makefile: call protoc library, which is responsible for compiling the protobuf definition into code. In addition usage of gRPC is specified as well as context and ouput path <br/> 
 
