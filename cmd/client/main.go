package main

import (
	"flag"
	"fmt"
	"os"
	gr "test_kafka/pkg/client"
	lg "test_kafka/pkg/logger"
	kp "test_kafka/pkg/producer"
)

func main() {
	//fmt.Printf("Client init \n")
	// Load init parameters
	grpcHostName := flag.String("grpchost", "", "grpc host name")
	grpcPort := flag.String("grpcport", "8086", "port for grpc connect")

	kafkaHostName := flag.String("kafkahost", "localhost", "kafka host name")
	kafkaPort := flag.String("kafkaport", "9092", "port for kafka connect")

	logLevel := flag.String("loglvl", "", "logging message level")
	logFile := flag.String("logfile", "", "logging message to file")
	flag.Parse()

	//Init log system
	log := lg.LogInit(*logLevel, *logFile)
	log.Println("Client log system init.")
	lg.PrintOsArgs(log)

	// Check existing obligatory http and db parameters
	if *grpcHostName == "" {
		flag.PrintDefaults()
		fmt.Println("")
		log.Fatalln("Init client error: set not all obligatory parameters.")
		os.Exit(1)
	}

	//Init kafka producer
	p := kp.New(*kafkaHostName, *kafkaPort, "Test", log)
	//Start gRPC cient
	g := gr.New(*grpcHostName, *grpcPort, log)
	err := g.Start(p)
	if err != nil {
		log.Fatalf("gRPC error: %s.", err.Error())
	}
}
