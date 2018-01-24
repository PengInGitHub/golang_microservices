//conignment-service/main.go

//make some changes

package main

import(
	"log"
	"net"
	//import the generated protobuf code
	pb "github.com/ewanvalentine/shipper/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	"google/golang.org/grpc"
	"google/golang.org/grpc/reflection"
)

const(
	port = ":50051"
)

type IRepository interface{
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct{
	consignment []*pb.Consignment
}

func (repo *Repository) Create(consigment *pb.Consignment )(*pb.Consignment, error){
	updated := append(repo.consignments, consignment)
	repo.consignments = updated 
	return consignment, nil
}

