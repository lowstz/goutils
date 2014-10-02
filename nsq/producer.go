package nsq

import (
	"github.com/bitly/go-nsq"
)

var (
	NsqProducer *nsq.Producer
	NsqAddress  string
)

func GetProducr(host string) (*nsq.Producer, error) {
	//NsqAddress = "127.0.0.1:4150"
	NsqAddress = host
	if NsqProducer != nil {
		return NsqProducer, nil
	}
	config := nsq.NewConfig()
	w, err := nsq.NewProducer(NsqAddress, config)
	NsqProducer = w
	return w, err
}
