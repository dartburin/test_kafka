package server

import (
	"context"
	"fmt"
	"net"

	se "test_kafka/api/proto"

	mid "test_kafka/pkg/middleware"

	kc "test_kafka/pkg/consumer"

	lg "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

// Data for handlers
type supportGRPC struct {
	log   lg.FieldLogger
	GPort string
	GConn string
}

// New creates new server struct
func New(gport string, log lg.FieldLogger) *supportGRPC {
	gstr := fmt.Sprintf(":%s", gport)
	log.Printf("Server gRPC init %s.", gstr)
	return &supportGRPC{
		log:   log,
		GPort: gport,
		GConn: gstr,
	}
}

func (s *supportGRPC) Start(cons *kc.CGHandler) error {
	//Consumer
	go cons.ConsumerGroup()

	s.log.Println("Server gRPC init.")
	lis, err := net.Listen("tcp", s.GConn)
	if err != nil {
		return err
	}

	ll := &mid.LogData{
		Log:  s.log,
		Name: "Server gRPC",
	}

	srv := grpc.NewServer(ll.WithServerUnaryInterceptors())
	se.RegisterSenderServer(srv, s)

	s.log.Println("Server gRPC start.")
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// SendMessage - handler for request
func (s *supportGRPC) SendMessage(ctx context.Context, imsg *se.SendMsg) (*se.Result, error) {
	s.log.Println("Server gRPC get message.")
	bb := se.Result{}
	return &bb, nil
}
