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

