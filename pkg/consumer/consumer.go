package consumer

import (
	"context"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	lg "github.com/sirupsen/logrus"
)

type CGHandler struct {
	log   lg.FieldLogger
	CHost string
	CPort string
	CConn string
	Topic string
}

// New creates new server struct
func New(chost string, cport string, ctopic string, log lg.FieldLogger) *CGHandler {
	cstr := fmt.Sprintf("%s:%s", chost, cport)
	log.Printf("Consumer init %s.", cstr)
	return &CGHandler{
		log:   log,
		CHost: chost,
		CPort: cport,
		CConn: cstr,
		Topic: ctopic,
	}
}

func (*CGHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (*CGHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h *CGHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func (h *CGHandler) ConsumerGroup() {

	config := sarama.NewConfig()
	config.Version = sarama.V2_7_0_0 // specify appropriate version
	config.Consumer.Return.Errors = true

	// Start with a client
	client, err := sarama.NewClient(strings.Split(h.CConn, ","), config)
	if err != nil {
		h.log.Printf("Consumer new client error: %v.", err)

	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient("Group_001", client)
	if err != nil {
		h.log.Printf("Consumer new group error: %v.", err)
	}
	defer func() { _ = group.Close() }()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := strings.Split(h.Topic, ",")

		err := group.Consume(ctx, topics, h)
		if err != nil {
			h.log.Printf("Consume error: %v.", err)
		}
	}
}
