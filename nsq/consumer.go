package nsq

import (
	"github.com/bitly/go-nsq"
)

type ConsumerQueue struct {
	Consumers       []*nsq.Consumer
	NsqdLookupdAddr string
}

type Handler func(*nsq.Message)

type Queue struct {
	Callback Handler
	*nsq.Consumer
}

func (q *Queue) HandleMessage(message *nsq.Message) error {
	q.Callback(message)
	return nil
}

func (c *ConsumerQueue) Register(topic string, channel string, handler Handler) {
	config := nsq.NewConfig()
	w, _ := nsq.NewConsumer(topic, channel, config)

	q := &Queue{
		handler, w,
	}

	w.AddConcurrentHandlers(q, 100)

	c.Consumers = append(c.Consumers, w)
}

func (c *ConsumerQueue) Connect() {
	for _, v := range c.Consumers {
		v.ConnectToNSQLookupd(c.NsqdLookupdAddr)
	}
}

func (c *ConsumerQueue) Start() {
	<-make(chan bool)

}
