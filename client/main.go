/*
	A simple test gRPC client
	go run client/main.go
*/

package main

import (
	"context"
	"fmt"
	"log"
	"time"
	pb "xevo/physeter-context-server/proto"

	google_protobuf "github.com/gogo/protobuf/types"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
)

const (
	address = "localhost:10000"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewContextServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t0, _ := time.Parse(time.RFC3339, "2018-05-01T12:30:00+09:00")
	t, _ := ptypes.TimestampProto(t0)

	request := &pb.GetRecommendsRequest{
		UserState: &pb.UserState{},
		CarState: &pb.CarState{
			CurrentLocation:     &pb.Coordinates{Latitude: 35.650645, Longitude: 139.710048},
			FuelLevelPercentage: 30,
		},
		Time: &google_protobuf.Timestamp{Seconds: t.Seconds, Nanos: t.Nanos},
	}

	r, err := c.GetRecommends(ctx, request)
	if err != nil {
		log.Fatalf("could not search: %v", err)
	}
	for _, v := range r.Recommends {
		fmt.Printf("name: %s\n", v.Name)
	}
}
