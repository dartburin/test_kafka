# test_kafka
The test for Kafka with gRPC.


Test method
-----------

Begin
.....

(docker terminal) docker-compose up --build

(client terminal) 
1. For local client:
Compile client:
	go build cmd/client/main.go
Run client:
	./main -grpchost=localhost -grpcport=8086
	./main -grpchost=localhost -grpcport=8086 -kafkaport=49154


OR
2. For client in the client container:
Start bash in the container:
	docker exec -it gRPCclient bash
Run client:
	./main -grpchost=serv -grpcport=8086
	./main -grpchost=serv -grpcport=8086 -kafkahost=172.17.0.1 -kafkaport=49154

Note: If you will quit from client programm, please, type ~QUIT and press ENTER.


End
...

(docker terminal) Ctrl+z
(docker terminal) docker-compose down


View console output:
--------------------
(log terminal) docker logs -f gRPCserver

For ProtoBuf generation:
------------------------ 
protoc -I. -I$GOPATH/src --go_out=plugins=grpc:. api/proto/sender.proto
