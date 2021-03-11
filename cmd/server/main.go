package main

import (
	"flag"
	"fmt"
	"os"
	lg "test_kafka/pkg/logger"

	kc "test_kafka/pkg/consumer"
	gp "test_kafka/pkg/server"
)

func main() {
	// Load init parameters
	logLevel := flag.String("loglvl", "", "logging message level")
	logFile := flag.String("logfile", "", "logging message to file")
	grpcPort := flag.String("grpcport", "", "port for grpc connect")

	kafkaHostName := flag.String("kafkahost", "localhost", "kafka host name")
	kafkaPort := flag.String("kafkaport", "9092", "port for kafka connect")

	flag.Parse()

	//Init log system
	log := lg.LogInit(*logLevel, *logFile)
	log.Println("Server log system init.")
	lg.PrintOsArgs(log)

	// Check existing obligatory http and db parameters
	if *grpcPort == "" {
		flag.PrintDefaults()
		fmt.Println("")
		log.Fatalln("Init server error: set not all obligatory parameters.")
		os.Exit(1)
	}

	//Init kafka consumer
	cons := kc.New(*kafkaHostName, *kafkaPort, "Test", log)
	//Start gRPC server
	g := gp.New(*grpcPort, log)
	err := g.Start(cons)
	if err != nil {
		log.Fatalf("Server gRPC error: %s.", err.Error())
	}
}
