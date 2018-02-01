//consignment-cli/cli.go
package main
import(
    "encoding/json"
    "io/ioutil"
    "log"
    "os"

    pb "github.com/user/golang_microservice/ewanvalentine/shipper/consignment-service/proto/consignment"
    "golang/org/x/net/context"
    "google.golang.org/grpc"
)

const(
    address = "localhost:500501"
    defaultFilename = "consignment.json"
)

func parseFile(file string)(*pb.Consignment, error){
    var consignment *pb.Consignment
    data, err := ioutil.ReadFile(file)
    if err != nil{
        return nil, err
    }
    json.Unmarshal(data,&consignment)
    return consignment, err
}

func main{
    //set up a connection to the server
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil{
        log.Fatalf("Could not greet: %v ", err)
    }
    log.Printf("Created: %t", r.Created)
}