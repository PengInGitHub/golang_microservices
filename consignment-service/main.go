//consignment-service/main.go
package main

import (
    "log"
    "net"
    //Import the generated protobuf code
    pb "github.com/user/golang_microservice/ewanvalentine/shipper/consignment-service/proto/consignment"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

const(
    port = ":50051"
)

type IRepository interface{
    Create(*pb.Consignment)(*pb.Consignment, error) 
}

//Repository - Dummy repository, this simulates a use of datastore of some kind.
//We'll replace this with a real implementation later on

type Repository struct{
    consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment)(*pb.Consignment, error){
    updated := append(repo.consignments, consignment)
    repo.consignments = updated
    return consignment, nil 
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc 
// to give you a better idea. 

type service struct{
    repo IRepositry
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an 
// argument, these are handled by the gRPC server. 

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment)(*pb.Response, error){
    //save our consignment
    consignment, err := s.repo.Create(req)
    if err != nil {
        return nil, err
    }
    //Return matching the 'Response' message we created in our protobuf definition
    return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main(){
    repo := Repositry{}
    
    //set-up our grpc server
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    // Register our service with the gRPC server, this will tie our
    // implementation into the auto-generated interface code for our
    // protobuf definition. 
    pb.RegisterShippingServiceServer(s, &service{repo})
    
    //Register reflection service on gRPC server
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}




















