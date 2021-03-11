package producer

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	lg "github.com/sirupsen/logrus"
)

// Data for gRPC clent
type ProducerKafka struct {
	log   lg.FieldLogger
	Host  string
	KPort string
	KConn string
	Topic string
}

// New creates the new Kafka producer struct
func New(khost string, kport string, ktopic string, log lg.FieldLogger) *ProducerKafka {
	kstr := fmt.Sprintf("%s:%s", khost, kport)
	return &ProducerKafka{
		log:   log,
		Host:  khost,
		KPort: kport,
		KConn: kstr,
		Topic: ktopic,
	}
}

// setupProducer will create a AsyncProducer and returns it
func (p *ProducerKafka) SetupProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	return sarama.NewAsyncProducer(strings.Split(p.KConn, ","), config)
}

// produceMessage will send str
func (p *ProducerKafka) ProduceMessage(producer sarama.AsyncProducer, str string) {
	message := &sarama.ProducerMessage{Topic: p.Topic, Value: sarama.StringEncoder(str)}
	producer.Input() <- message
}
