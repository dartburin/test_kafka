package client

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	cod "test_kafka/api/encode"
	se "test_kafka/api/proto"

	kp "test_kafka/pkg/producer"

	lg "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Data for gRPC clent
type clientGRPC struct {
	log   lg.FieldLogger
	Host  string
	GPort string
	GConn string
}

// New creates the new gRPC clent struct
func New(ghost string, gport string, log lg.FieldLogger) *clientGRPC {
	gstr := fmt.Sprintf("%s:%s", ghost, gport)
	return &clientGRPC{
		log:   log,
		Host:  ghost,
		GPort: gport,
		GConn: gstr,
	}
}

// Start gRPC client
func (s *clientGRPC) Start(prod *kp.ProducerKafka) error {
	cmdQuit := "~QUIT"

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(s.GConn, opts...)

	if err != nil {
		s.log.Fatalf(" connect error: %v", err)
	}
	s.log.Println("gRPC client connect.")
	defer conn.Close()

	client := se.NewSenderClient(conn)

	conReader := bufio.NewReader(os.Stdin)

	s.log.Println("gRPC client start.")

	//producer
	aprod, err := prod.SetupProducer()
	if err != nil {
		s.log.Fatalf("producer error: %v", err)
	}

	for {
		input, err := conReader.ReadString('\n')
		if err != nil {
			continue
		}
		str := strings.TrimSuffix(input, "\n")
		fmt.Printf("Client read: %s\n", str)

		if str == cmdQuit {
			fmt.Printf("\nClient: Quit\n")
			break
		} else {
			buff := cod.Encode(str)

			prod.ProduceMessage(aprod, buff.String())

			msg := &se.SendMsg{
				Msg: buff.Bytes(),
			}
			response, err := client.SendMessage(context.Background(), msg)

			if err != nil {
				s.log.Printf("fail to dial: %v", err)
			}

			fmt.Printf("Client: %v\n", response)
		}
	}

	return nil
}
